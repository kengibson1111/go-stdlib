package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	names := []string{"bob", "Bob", "alice", "Vera", "VERA"}
	fmt.Println(names)

	names = slices.CompactFunc(names, strings.EqualFold)
	fmt.Println(names)
}
