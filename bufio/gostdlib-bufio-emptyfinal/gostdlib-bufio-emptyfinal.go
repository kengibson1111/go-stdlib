package main

import (
	"bufio"
	"log"
	"strings"
)

func main() {
	// Comma-separated list; last entry is empty.
	const input = "1,2,3,4,"

	scanner := bufio.NewScanner(strings.NewReader(input))
	if scanner == nil {
		log.Fatal("bufio.NewScanner(): nil scanner")
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

	// Set the split function for the scanning operation.
	scanner.Split(onComma)
	if err := scanner.Err(); err != nil {
		log.Fatal("scanner.Split():", err)
	}

	// Scan. Error handling behavior in the loop
	// is to log the error in the application log and then break. The error will
	// also be logged as an error using log.Fatal(). log.Fatal() will also
	// terminate the process, so it should only be used in a main() function.
	//
	// The use of log.Print() and log.Println() within the loop in conjunction with
	// log.Fatal() after the loop is a way to tell a full story about the function
	// flow without changing the actual error. The application log will show the
	// error that occurred in context of the loop, and it will show the error after
	// the loop exit. It tells the full story instead of the story being split
	// between the application log and the system error log.
	//
	// In this example, the default application log is os.Stdout which is fine for
	// the example. In a real deploy, the application log target would need to be
	// configured.
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Println("scanner.Scan():", err)
			break
		}

		text := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Println("scanner.Text():", err)
			break
		}

		log.Printf("%q", text)
	}

	// notice how this block can be hit as a result of an error that propagated up from scanner.Scan().
	// Ideally, scanner.Scan() would handle all errors gracefully and allow the first error check in the
	// for loop to be used for any scanner.Scan() error regardless of the error source.
	if err := scanner.Err(); err != nil {
		log.Fatal("reading input:", err)
	}
}
