package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("TrailingZeros8(%08b) = %d\n", 14, bits.TrailingZeros8(14))
	fmt.Printf("TrailingZeros16(%016b) = %d\n", 14, bits.TrailingZeros16(14))
	fmt.Printf("TrailingZeros32(%032b) = %d\n", 14, bits.TrailingZeros32(14))
	fmt.Printf("TrailingZeros64(%064b) = %d\n", 14, bits.TrailingZeros64(14))
}
