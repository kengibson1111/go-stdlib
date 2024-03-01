package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	_, err := s.Seek(10, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
}
