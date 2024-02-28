package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	buf := make([]byte, 9)

	_, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buf)
}
