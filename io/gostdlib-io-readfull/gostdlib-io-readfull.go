package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)

	_, err := io.ReadFull(r, buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)

	// this continues reading from the same reader.
	_, err = io.ReadFull(r, longBuf)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%s\n", longBuf)
}
