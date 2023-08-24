package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"
)

func main() {
	buff := []byte{31, 139, 8, 0, 0, 9, 110, 136, 0, 255, 202, 72, 205, 201, 201, 215, 81, 40, 207, 47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 83, 116, 36, 244, 13, 0, 0, 0}
	b := bytes.NewReader(buff)

	r, err := gzip.NewReader(b)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r)

	r.Close()
}
