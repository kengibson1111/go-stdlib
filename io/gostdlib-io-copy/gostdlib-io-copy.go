package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}
