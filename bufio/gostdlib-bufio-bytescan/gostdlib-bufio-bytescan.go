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

		byteArray := scanner.Bytes()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "scanner.Bytes():", err)
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
}
