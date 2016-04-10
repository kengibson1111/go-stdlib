package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

func main() {
	c := 10
	b := make([]byte, c)
	fmt.Printf("Before: %x\n", b)

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("After: %x\n", b)

	// The slice should now contain random bytes instead of only zeroes.
	fmt.Println(bytes.Equal(b, make([]byte, c)))

}
