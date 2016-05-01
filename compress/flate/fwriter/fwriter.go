package main

import (
	"bytes"
	"compress/flate"
	"fmt"
)

func main() {
	var b bytes.Buffer

	w, err := flate.NewWriter(&b, flate.DefaultCompression)
	if err != nil {
		panic(err)
	}

	w.Write([]byte("hello, world\n"))
	w.Close()
	fmt.Println(b.Bytes())
}
