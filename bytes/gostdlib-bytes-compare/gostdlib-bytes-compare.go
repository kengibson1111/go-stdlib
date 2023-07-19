package main

import (
	"bytes"
	"fmt"
)

func compareEqual(a, b []byte) {
	if bytes.Compare(a, b) < 0 {
		fmt.Println("a < b")
	}
	if bytes.Compare(a, b) <= 0 {
		fmt.Println("a <= b")
	}
	if bytes.Compare(a, b) > 0 {
		fmt.Println("a > b")
	}
	if bytes.Compare(a, b) >= 0 {
		fmt.Println("a >= b")
	}

	// Prefer Equal to Compare for equality comparisons.
	if bytes.Equal(a, b) {
		fmt.Println("a == b")
	}
	if !bytes.Equal(a, b) {
		fmt.Println("a != b")
	}
}

func main() {
	// Interpret Compare's result by comparing it to zero.
	var a, b []byte
	fmt.Printf("\na = %s, b = %s\n", a, b)
	compareEqual(a, b)

	a = []byte("a")
	b = []byte("b")
	fmt.Printf("\na = %s, b = %s\n", a, b)
	compareEqual(a, b)

	a = []byte("a")
	b = []byte("A")
	fmt.Printf("\na = %s, b = %s\n", a, b)
	compareEqual(a, b)

	a = []byte("a")
	b = []byte("aa")
	fmt.Printf("\na = %s, b = %s\n", a, b)
	compareEqual(a, b)
}
