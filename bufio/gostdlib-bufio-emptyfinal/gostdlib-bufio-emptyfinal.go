package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Comma-separated list; last entry is empty.
	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
	if scanner == nil {
		fmt.Fprintln(os.Stderr, "bufio.NewScanner(): nil scanner")
		return
	}

	// Define a split function that separates on commas.
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}

		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}

	scanner.Split(onComma)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scanner.Split():", err)
		return
	}

	// Scan.
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

		fmt.Printf("%q ", text)
	}

	fmt.Println()

	// notice how this block can be hit as a result of an error that propagated up from scanner.Scan().
	// Ideally, scanner.Scan() would handle all errors gracefully and allow the first error check in the
	// for loop to be used for any scanner.Scan() error regardless of the error source.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
