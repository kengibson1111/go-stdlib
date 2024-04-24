package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("Len8(%08b) = %d\n", 8, bits.Len8(8))
	fmt.Printf("Len16(%016b) = %d\n", 8, bits.Len16(8))
	fmt.Printf("Len32(%032b) = %d\n", 8, bits.Len32(8))
	fmt.Printf("Len64(%064b) = %d\n", 8, bits.Len64(8))
}
