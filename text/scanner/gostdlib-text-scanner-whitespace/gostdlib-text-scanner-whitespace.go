package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	// tab-separated values
	const src = `aa	ab	ac	ad
ba	bb	bc	bd
ca	cb	cc	cd
da	db	dc	dd`

	var (
		col, row int
		s        scanner.Scanner
		tsv      [4][4]string // large enough for example above
	)
	s.Init(strings.NewReader(src))
	s.Whitespace ^= 1<<'\t' | 1<<'\n' // don't skip tabs and new lines

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case '\n':
			row++
			col = 0
		case '\t':
			col++
		default:
			tsv[row][col] = s.TokenText()
		}
	}

	fmt.Println(tsv)
}
