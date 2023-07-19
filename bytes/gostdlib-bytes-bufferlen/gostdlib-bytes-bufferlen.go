package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Grow(64)

	b.Write([]byte("abcde"))
	fmt.Printf("%d %d\n", b.Len(), b.Cap())

	b.WriteTo(os.Stdout)
	fmt.Printf("\n%d %d\n", b.Len(), b.Cap())
}
