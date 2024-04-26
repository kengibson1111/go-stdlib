package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("OnesCount(%b) = %d\n", 14, bits.OnesCount(14))
	fmt.Printf("OnesCount8(%08b) = %d\n", 14, bits.OnesCount8(14))
	fmt.Printf("OnesCount16(%016b) = %d\n", 14, bits.OnesCount16(14))
	fmt.Printf("OnesCount32(%032b) = %d\n", 14, bits.OnesCount32(14))
	fmt.Printf("OnesCount64(%064b) = %d\n", 14, bits.OnesCount64(14))
}
