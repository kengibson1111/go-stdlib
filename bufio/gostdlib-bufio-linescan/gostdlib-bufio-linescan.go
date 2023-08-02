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
}
