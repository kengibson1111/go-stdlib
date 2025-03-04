package main

import (
	"fmt"
	"slices"
)

func main() {
	seq := []int{0, 1, 1, 2, 3, 5, 8}

	seq = slices.DeleteFunc(seq, func(n int) bool {
		return n%2 != 0 // delete the odd numbers
	})

	fmt.Println(seq)
}
