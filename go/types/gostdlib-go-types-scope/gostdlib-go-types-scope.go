package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"regexp"
	"strings"
)

func main() {
	// Parse the source files for a package.
	fset := token.NewFileSet()
	var files []*ast.File
	for _, src := range []string{
		`package main
import "fmt"
func main() {
	freezing := FToC(-18)
	fmt.Println(freezing, Boiling) }
`,
		`package main
import "fmt"
type Celsius float64
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func FToC(f float64) Celsius { return Celsius(f - 32 / 9 * 5) }
const Boiling Celsius = 100
func Unused() { {}; {{ var x int; _ = x }} } // make sure empty block scopes get printed
`,
	} {
		files = append(files, mustParse(fset, src))
	}

	// Type-check a package consisting of these files.
	// Type information for the imported "fmt" package
	// comes from $GOROOT/pkg/$GOOS_$GOOARCH/fmt.a.
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("temperature", fset, files, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print the tree of scopes.
	// For determinism, we redact addresses.
	var buf strings.Builder
	pkg.Scope().WriteTo(&buf, 0, true)
	rx := regexp.MustCompile(` 0x[a-fA-F\d]*`)
	fmt.Println(rx.ReplaceAllString(buf.String(), ""))
}

func mustParse(fset *token.FileSet, src string) *ast.File {
	f, err := parser.ParseFile(fset, pkgName(src), src, 0)

	if err != nil {
		panic(err) // so we don't need to pass *testing.T
	}

	return f
}

func pkgName(src string) string {
	const kw = "package "

	if i := strings.Index(src, kw); i >= 0 {
		after := src[i+len(kw):]
		n := len(after)

		if i := strings.IndexAny(after, "\n\t ;/"); i >= 0 {
			n = i
		}

		return after[:n]
	}

	panic("missing package header: " + src)
}
