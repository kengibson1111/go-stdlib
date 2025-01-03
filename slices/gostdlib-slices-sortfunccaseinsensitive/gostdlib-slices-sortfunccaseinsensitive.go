package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	names := []string{"Bob", "alice", "VERA"}
	fmt.Println(names)

	slices.SortFunc(names, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})

	fmt.Println(names)
}
