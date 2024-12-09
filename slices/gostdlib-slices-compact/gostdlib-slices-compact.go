package main

import (
	"fmt"
	"slices"
)

func main() {
	seq := []int{0, 1, 1, 2, 3, 5, 8}
	fmt.Println(seq)

	seq = slices.Compact(seq)
	fmt.Println(seq)
}
