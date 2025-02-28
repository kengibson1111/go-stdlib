package main

import (
	"fmt"
	"strings"
)

func main() {
	f := func(r rune) bool {
		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
	}

	fmt.Println(strings.ContainsFunc("hello", f))
	fmt.Println(strings.ContainsFunc("rhythms", f))
}
