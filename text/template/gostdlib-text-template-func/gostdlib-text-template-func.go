package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		// The name "upper" is what the function will be called in the template text.
		"upper": strings.ToUpper,
	}

	// A simple template definition to test our function.
	// We print the input text several ways:
	// - the original
	// - upper-cased
	// - upper-cased and then printed with %q
	// - printed with %q and then upper-cased.
	const templateText = `
Input: {{printf "%q" .}}
Output 0: {{upper .}}
Output 1: {{upper . | printf "%q"}}
Output 2: {{printf "%q" . | upper}}
`

	// Create a template, add the function map, and parse the text.
	tmpl, err := template.New("upperTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Run the template to verify the output.
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}
