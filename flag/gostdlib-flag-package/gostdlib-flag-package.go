package main

import (
	"flag"
	"fmt"
)

// Example 1: A single string flag called "species" with default value "gopher".
var species = flag.String("species", "gopher", "the species we are studying")

func main() {
	flag.Parse()
	flag.Usage()

	fmt.Println()
	fmt.Println("species has value ", *species)
}
