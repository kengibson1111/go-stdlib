package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read")

	_, err := io.CopyN(os.Stdout, r, 4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
}
