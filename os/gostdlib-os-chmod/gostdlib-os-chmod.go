package main

import (
	"log"
	"os"
)

func main() {
	if err := os.Chmod("temp.txt", 0644); err != nil {
		log.Fatal(err)
	}
}
