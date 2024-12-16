package main

import (
	"fmt"
	"slices"
)

func main() {
	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Println(letters)

	letters = slices.Delete(letters, 1, 4)
	fmt.Println(letters)
}
