package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera", "Zac"}
	fmt.Println(names)

	names = slices.Replace(names, 1, 3, "Bill", "Billie", "Cat")
	fmt.Println(names)
}
