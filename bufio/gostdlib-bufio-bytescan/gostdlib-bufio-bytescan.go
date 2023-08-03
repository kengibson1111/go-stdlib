package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner == nil {
		fmt.Fprintln(os.Stderr, "bufio.NewScanner(): nil scanner")
		return
	}

	// default split function is ScanLines()
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "scanner.Scan():", err)
			break
		}

		byteArray := scanner.Bytes()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "scanner.Bytes():", err)
			break
		}

		fmt.Println("len =", len(byteArray))
		if len(byteArray) == 0 {
			// break out if an empty line
			break
		}

		for i := 0; i < len(byteArray); i++ {
			fmt.Print(byteArray[i])
			if i < len(byteArray)-1 {
				fmt.Print(",")
			}
		}

		fmt.Println()
	}

	// notice how this block can be hit as a result of an error that propagated up from scanner.Scan().
	// Ideally, scanner.Scan() would handle all errors gracefully and allow the first error check in the
	// for loop to be used for any scanner.Scan() error regardless of the error source.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid input: ", err)
	}
}
