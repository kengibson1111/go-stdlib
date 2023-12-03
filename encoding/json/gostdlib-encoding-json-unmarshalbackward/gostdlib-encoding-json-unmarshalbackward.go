package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// the blob is version 1 of an event.
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)

	// this is a version 2 type definition
	type Animal2 struct {
		Name     string
		Order    string
		NewOrder string
	}

	// this is a version 3 type definition
	type Animal3 struct {
		Name     string
		NewOrder string
	}

	var animals2 []Animal2
	err := json.Unmarshal(jsonBlob, &animals2)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v\n", animals2)

	var animals3 []Animal3
	err = json.Unmarshal(jsonBlob, &animals3)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v\n", animals3)
}
