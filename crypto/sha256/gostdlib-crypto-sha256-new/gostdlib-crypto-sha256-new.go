package main

import (
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	h := sha256.New()
	io.WriteString(h, "His money is twice tainted:")
	io.WriteString(h, " 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
	fmt.Println()

	data := []byte("His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", sha256.Sum256(data))
	fmt.Println()

	h224 := sha256.New224()
	io.WriteString(h224, "His money is twice tainted:")
	io.WriteString(h224, " 'taint yours and 'taint mine.")
	fmt.Printf("% x", h224.Sum(nil))
	fmt.Println()

	data224 := []byte("His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", sha256.Sum224(data224))
	fmt.Println()
}
