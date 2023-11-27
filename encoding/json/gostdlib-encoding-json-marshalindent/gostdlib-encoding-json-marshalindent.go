package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.MarshalIndent(group, "", "\t") // 2nd param is a prefix that starts each newline that has indentation.
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)
	fmt.Println()
}
