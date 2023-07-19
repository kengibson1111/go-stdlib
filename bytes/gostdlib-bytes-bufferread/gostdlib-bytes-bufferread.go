package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Grow(64)

	b.Write([]byte("abcde"))

	fmt.Printf("%d %s %d %d\n", b.Len(), string(b.Next(2)), b.Len(), b.Cap())
	fmt.Printf("%d %s %d %d\n", b.Len(), string(b.Next(2)), b.Len(), b.Cap())
	fmt.Printf("%d %s %d %d\n", b.Len(), string(b.Next(2)), b.Len(), b.Cap())
}
