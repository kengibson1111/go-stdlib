package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	// like the basic sample.
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// 1st half of byte 7 is "4". So the first OR guarantees
	// that the 1st half of byte 7 is at least 4. The AND guarantees that the
	// max value of the 1st half of byte 7 is 4. Always 4.
	b[6] = (b[6] | 0x40) & 0x4F
	// 1st half of byte 9 is "8", "9", "a", or "b". The first
	// OR guarantees that the 1st half of byte 9 is at least 8. The AND
	// guarantees that the max value of the 1st half of byte 9 is b.
	b[8] = (b[8] | 0x80) & 0xBF
	
	//done
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Println("UUID: ", uuid)
}
