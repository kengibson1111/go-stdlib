package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// this make sure that the 13th character is "4"
	b[6] = (b[6] | 0x40) & 0x4F
	// this make sure that the 17th is "8", "9", "a", or "b"
	b[8] = (b[8] | 0x80) & 0xBF

	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Println("UUID: ", uuid)
}
