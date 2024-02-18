package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	_, err := io.CopyBuffer(os.Stdout, r1, buf)
	if err != nil {
		log.Fatal(err)
	}

	// ... reused here also. No need to allocate an extra buffer.
	_, err = io.CopyBuffer(os.Stdout, r2, buf)
	if err != nil {
		log.Fatal(err)
	}
}
