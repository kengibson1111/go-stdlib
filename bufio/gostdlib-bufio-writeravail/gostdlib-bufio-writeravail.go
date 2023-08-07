package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	if writer == nil {
		fmt.Fprintln(os.Stderr, "bufio.NewWriter(): nil writer")
		return
	}

	fmt.Println("writer buffer default size =", writer.Available())

	for _, i := range []int64{1, 2, 3, 4} {
		b := writer.AvailableBuffer()
		fmt.Printf("append %d before = %q %d %d\n", i, b, len(b), cap(b))

		b = strconv.AppendInt(b, i, 10)
		b = append(b, ' ')
		fmt.Printf("append %d after = %q %d %d\n", i, b, len(b), cap(b))

		writer.Write(b)
		fmt.Printf("write %d after = %q %d %d\n", i, b, len(b), cap(b))

		// anything after a Write and before reassigning the AvailableBuffer is lost.
		b = append(b, ' ')
		fmt.Printf("invalid append %d after write = %q %d %d\n", i, b, len(b), cap(b))
	}

	err := writer.Flush() // Don't forget to flush!
	if err != nil {
		fmt.Fprintln(os.Stderr, "writer.Flush():", err)
	}

	fmt.Println()
}
