package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	if writer == nil {
		log.Fatal("bufio.NewWriter(): nil writer")
	}

	log.Println("writer buffer default size =", writer.Available())

	for _, i := range []int64{1, 2, 3, 4} {
		b := writer.AvailableBuffer()
		log.Printf("append %d before = %q %d %d\n", i, b, len(b), cap(b))

		b = strconv.AppendInt(b, i, 10)
		b = append(b, ' ')
		log.Printf("append %d after = %q %d %d\n", i, b, len(b), cap(b))

		writer.Write(b)
		log.Printf("write %d after = %q %d %d\n", i, b, len(b), cap(b))

		// anything after a Write and before reassigning the AvailableBuffer is lost.
		b = append(b, ' ')
		log.Printf("invalid append %d after write = %q %d %d\n", i, b, len(b), cap(b))
	}

	err := writer.Flush() // Don't forget to flush!
	if err != nil {
		log.Fatal("writer.Flush():", err)
	}

	fmt.Println()
}
