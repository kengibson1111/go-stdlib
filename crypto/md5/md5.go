package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "His money is twice tainted:")
	io.WriteString(h, " 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
	fmt.Println()

	data := []byte("His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", md5.Sum(data))
	fmt.Println()
}
