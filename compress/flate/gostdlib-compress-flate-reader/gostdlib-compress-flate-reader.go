package main

import (
	"bytes"
	"compress/flate"
	"io"
	"os"
)

func main() {
	buff := []byte{202, 72, 205, 201, 201, 215, 81, 40, 207, 47, 202, 73, 225, 2, 4, 0, 0, 255, 255}
	b := bytes.NewReader(buff)

	r := flate.NewReader(b)
	io.Copy(os.Stdout, r)
	r.Close()
}
