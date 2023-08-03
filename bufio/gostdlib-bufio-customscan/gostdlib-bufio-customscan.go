package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// An artificial input source.
	const input = "1234 5678 1234567901234567890"

	scanner := bufio.NewScanner(strings.NewReader(input))
	if scanner == nil {
		fmt.Fprintln(os.Stderr, "bufio.NewScanner(): nil scanner")
		return
	}

	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}

		return
	}

	// Set the split function for the scanning operation.
	scanner.Split(split)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scanner.Split():", err)
		return
	}

	// Validate the input
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "scanner.Scan():", err)
			break
		}

		text := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "scanner.Text():", err)
			break
		}

		fmt.Printf("%s\n", text)
	}

	// notice how this block can be hit as a result of an error that propagated up from scanner.Scan().
	// Ideally, scanner.Scan() would handle all errors gracefully and allow the first error check in the
	// for loop to be used for any scanner.Scan() error regardless of the error source.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid input: ", err)
	}
}
