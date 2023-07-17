package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader("gopher"))
	var byteArray []byte

	for scanner.Scan() {
		byteArray = scanner.Bytes()
		fmt.Println(len(byteArray) == 6)
	}

	for i := 0; i < len(byteArray); i++ {
		fmt.Print(byteArray[i])
		if i < len(byteArray)-1 {
			fmt.Print(",")
		}
	}

	fmt.Println()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}
}
