package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// New version of an event that contains a new field.
	var jsonBlob1 = []byte(`[
		{"Version": "2", "Name": "Platypus", "Order": "Monotremata", "NewOrder": "Monotremata"},
		{"Version": "2", "Name": "Quoll", "Order": "Dasyuromorphia", "NewOrder": "Dasyuromorphia"}
	]`)
	
	// New version of an event that contains a new field and doesn't contain 
	// a field from the previous version.
	var jsonBlob2 = []byte(`[
		{"Version": "3", "Name": "Platypus", "NewOrder": "Monotremata"},
		{"Version": "3", "Name": "Quoll",    "NewOrder": "Dasyuromorphia"}
	]`)
	
	// type represents the previous version of the event.
	type Animal struct {
		Version string
		Name  string
		Order string
	}
	
	var animals1 []Animal
	err := json.Unmarshal(jsonBlob1, &animals1)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals1)
	
	var animals2 []Animal
	err = json.Unmarshal(jsonBlob2, &animals2)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals2)
}
