package main

import (
	"bytes"
	"compress/lzw"
	"io"
	"os"
)

func main() {
	buff := []byte{104, 202, 176, 97, 243, 134, 5, 136, 59, 111, 228, 176, 33, 163, 32, 32}
	b := bytes.NewReader(buff)

	r := lzw.NewReader(b, lzw.LSB, 8)
	io.Copy(os.Stdout, r)
	r.Close()
}
