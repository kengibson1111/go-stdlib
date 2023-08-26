package main

import (
	"bytes"
	"compress/lzw"
	"fmt"
)

func main() {
	var b bytes.Buffer

	w := lzw.NewWriter(&b, lzw.LSB, 8)
	w.Write([]byte("hello, world\n"))
	w.Close()

	fmt.Println(b.Bytes())
}
