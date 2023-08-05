package main

import (
	"bytes"
	"fmt"
)

func main() {
	oink := []byte("oink oink oink")
	moo := bytes.ReplaceAll(oink, []byte("oink"), []byte("moo"))
	fmt.Printf("%s, %s\n", oink, moo)
}
