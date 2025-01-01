package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"alice", "Bob", "VERA"}
	fmt.Println(names)

	slices.Reverse(names)
	fmt.Println(names)
}
