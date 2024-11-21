package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("T")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1W")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("${1}W")))

	re2 := regexp.MustCompile(`a(?P<1W>x*)b`)
	fmt.Printf("%s\n", re2.ReplaceAll([]byte("-ab-axxb-"), []byte("$1W")))
	fmt.Printf("%s\n", re2.ReplaceAll([]byte("-ab-axxb-"), []byte("${1}W")))
}
