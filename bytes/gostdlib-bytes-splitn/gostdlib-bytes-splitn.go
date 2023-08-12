package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 2))
	fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), -1))

	z := bytes.SplitN([]byte(""), []byte(","), 1)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

	z2 := bytes.SplitN([]byte("a,b,c"), []byte(","), 0)
	fmt.Printf("%q (nil = %v)\n", z2, z2 == nil)
}
