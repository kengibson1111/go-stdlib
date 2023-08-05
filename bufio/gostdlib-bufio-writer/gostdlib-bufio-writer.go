package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	if writer == nil {
		fmt.Fprintln(os.Stderr, "bufio.NewWriter(): nil writer")
		return
	}

	fmt.Println("writer buffer default size =", writer.Available())

	fmt.Fprint(writer, "Hello, ")
	fmt.Fprint(writer, "world!\n")

	err := writer.Flush() // Don't forget to flush!
	if err != nil {
		fmt.Fprintln(os.Stderr, "writer.Flush():", err)
	}
}
