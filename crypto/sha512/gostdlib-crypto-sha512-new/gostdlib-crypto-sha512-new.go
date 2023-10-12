package main

import (
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	h := sha512.New()
	io.WriteString(h, "His money is twice tainted:")
	io.WriteString(h, " 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
	fmt.Println()

	data := []byte("His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", sha512.Sum512(data))
	fmt.Println()

	h384 := sha512.New384()
	io.WriteString(h384, "His money is twice tainted:")
	io.WriteString(h384, " 'taint yours and 'taint mine.")
	fmt.Printf("% x", h384.Sum(nil))
	fmt.Println()

	data384 := []byte("His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", sha512.Sum384(data384))
	fmt.Println()

	h224 := sha512.New512_224()
	io.WriteString(h224, "His money is twice tainted:")
	io.WriteString(h224, " 'taint yours and 'taint mine.")
	fmt.Printf("% x", h224.Sum(nil))
	fmt.Println()

	data224 := []byte("His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", sha512.Sum512_224(data224))
	fmt.Println()

	h256 := sha512.New512_256()
	io.WriteString(h256, "His money is twice tainted:")
	io.WriteString(h256, " 'taint yours and 'taint mine.")
	fmt.Printf("% x", h256.Sum(nil))
	fmt.Println()

	data256 := []byte("His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", sha512.Sum512_256(data256))
	fmt.Println()
}
