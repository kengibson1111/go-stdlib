package main

import (
	"log"
	"os"
)

func main() {
	err := os.WriteFile("testdir/hello.txt", []byte("Hello, Gophers!"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
