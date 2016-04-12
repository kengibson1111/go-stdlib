package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	h := sha1.New()
	io.WriteString(h, "His money is twice tainted:")
	io.WriteString(h, " 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
	fmt.Println()

	data := []byte("His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", sha1.Sum(data))
	fmt.Println()
}
