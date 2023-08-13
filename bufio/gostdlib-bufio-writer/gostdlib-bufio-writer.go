package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	if writer == nil {
		log.Fatal("bufio.NewWriter(): nil writer")
	}

	log.Println("writer buffer default size =", writer.Available())

	fmt.Fprint(writer, "Hello, ")
	fmt.Fprint(writer, "world!\n")

	err := writer.Flush() // Don't forget to flush!
	if err != nil {
		log.Fatal("writer.Flush():", err)
	}
}
