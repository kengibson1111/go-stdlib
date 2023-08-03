package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// default split function is ScanLines()
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "scanner.Scan():", err)
		}

		text := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "scanner.Text():", err)
		}

		fmt.Println("len =", len(text))
		if len(text) == 0 {
			// break out if an empty line
			break
		}

		fmt.Println(text)
	}

	// notice how this block can be hit as a result of an error that propagated up from scanner.Scan().
	// Ideally, scanner.Scan() would handle all errors gracefully and allow the first error check in the
	// for loop to be used for any scanner.Scan() error regardless of the error source.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid input: ", err)
	}
}
