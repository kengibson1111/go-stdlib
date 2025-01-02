package main

import (
	"fmt"
	"slices"
)

func main() {
	smallInts := []int8{0, 42, -10, 8}
	fmt.Println(smallInts)

	slices.Sort(smallInts)
	fmt.Println(smallInts)
}
