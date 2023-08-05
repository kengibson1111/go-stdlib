package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// An artificial input source.
	const input = "\tNow is the winter of our discontent,\n\tMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	if scanner == nil {
		fmt.Fprintln(os.Stderr, "bufio.NewScanner(): nil scanner")
		return
	}

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scanner.Split():", err)
		return
	}

	// Count the words.
	count := 0
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
		count++
	}

	fmt.Println()

	// notice how this block can be hit as a result of an error that propagated up from scanner.Scan().
	// Ideally, scanner.Scan() would handle all errors gracefully and allow the first error check in the
	// for loop to be used for any scanner.Scan() error regardless of the error source.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Printf("%d\n", count)
}
