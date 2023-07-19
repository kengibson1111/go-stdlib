package main

import (
	"bytes"
	"sort"
	"fmt"
)

func search(haystack [][]byte, needle []byte) {
	fmt.Printf("haystack = %v\n", haystack)
	fmt.Printf("needle = %v (%s)\n", needle, needle)

	i := sort.Search(len(haystack), func(i int) bool {
		// Return haystack[i] >= needle.
		return bytes.Compare(haystack[i], needle) >= 0
	})
	if i < len(haystack) && bytes.Equal(haystack[i], needle) {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}
}

func main() {
	// Binary search to find a matching byte slice.
	var needle []byte
	var haystack [][]byte
	search(haystack, needle)

	needle  = []byte("a")
	haystack = [][]byte{[]byte("a"), []byte("b"), []byte("c")}
	search(haystack, needle)

	needle  = []byte("c")
	search(haystack, needle)

	needle  = []byte("b")
	search(haystack, needle)

	needle  = []byte("d")
	search(haystack, needle)
}
