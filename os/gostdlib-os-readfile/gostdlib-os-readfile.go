package main

import (
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("testdir/testfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(data)
}
