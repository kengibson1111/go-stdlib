package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Grow(64)

	b.Write([]byte("abcde"))

	rdbuf := make([]byte, 1)
	n, err := b.Read(rdbuf)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
	fmt.Println(b.String())
	fmt.Println(string(rdbuf))
}
