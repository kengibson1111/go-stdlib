package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner == nil {
		log.Fatal("bufio.NewScanner(): nil scanner")
	}

	// default split function is ScanLines(). Error handling behavior in the loop
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

		byteArray := scanner.Bytes()
		if err := scanner.Err(); err != nil {
			log.Println("scanner.Bytes():", err)
			break
		}

		log.Println("len =", len(byteArray))
		if len(byteArray) == 0 {
			// break out if an empty line. This is expected behavior.
			break
		}

		log.Printf("hex = % x\n", byteArray)
		for i := 0; i < len(byteArray); i++ {
			log.Print(byteArray[i])
		}
	}

	// notice how this block can be hit as a result of an error that propagated up from scanner.Scan().
	// Ideally, scanner.Scan() would handle all errors gracefully and allow the first error check in the
	// for loop to be used for any scanner.Scan() error regardless of the error source.
	if err := scanner.Err(); err != nil {
		log.Fatal("Invalid input: ", err)
	}
}
