package main

import (
	"bytes"
	"fmt"
)

func main() {
	oink1 := []byte("oink")
	oinky1 := bytes.Replace(oink1, []byte("k"), []byte("ky"), 2)
	oink3 := []byte("oink oink oink")
	oinky3 := bytes.Replace(oink3, []byte("k"), []byte("ky"), 2)
	fmt.Printf("%s, %s, %s, %s\n", oink1, oinky1, oink3, oinky3)

	moo := bytes.Replace(oink3, []byte("oink"), []byte("moo"), -1)
	fmt.Printf("%s, %s\n", oink3, moo)
}
