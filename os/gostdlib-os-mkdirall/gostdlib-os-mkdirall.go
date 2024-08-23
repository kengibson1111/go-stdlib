package main

import (
	"log"
	"os"
)

func main() {
	err := os.MkdirAll("testdir/subdir", 0750)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("testdir/subdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}
