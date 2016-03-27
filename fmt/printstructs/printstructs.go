package main

import "fmt"

// An empty struct.
type EmptyStruct struct {}

// A struct with 6 fields.
type Struct1 struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}

func main() {
	var empty EmptyStruct
	var struct1 Struct1
	
	fmt.Printf("empty type = %T\n", empty)
	fmt.Printf("empty value = %v\n", empty)
	fmt.Printf("empty value with field names = %+v\n", empty)
	
	fmt.Printf("struct1 type = %T\n", struct1)
	fmt.Printf("struct1 value = %v\n", struct1)
	fmt.Printf("struct1 value with field names = %+v\n", struct1)
}
