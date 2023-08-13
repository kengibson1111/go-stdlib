package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func main() {
	// An artificial input source.
	const input = "1234 5678 1234567901234567890"

	scanner := bufio.NewScanner(strings.NewReader(input))
	if scanner == nil {
		log.Fatal("bufio.NewScanner(): nil scanner")
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
		log.Fatal("scanner.Split():", err)
	}

	// Validate the input. Error handling behavior in the loop
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

		log.Printf("%s\n", text)
	}

	// notice how this block can be hit as a result of an error that propagated up from scanner.Scan().
	// Ideally, scanner.Scan() would handle all errors gracefully and allow the first error check in the
	// for loop to be used for any scanner.Scan() error regardless of the error source.
	if err := scanner.Err(); err != nil {
		log.Fatal("Invalid input: ", err)
	}
}
