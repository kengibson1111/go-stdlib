package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%016b\n", 15)
	fmt.Printf("%016b\n", bits.ReverseBytes16(15))

	fmt.Printf("%032b\n", 15)
	fmt.Printf("%032b\n", bits.ReverseBytes32(15))

	fmt.Printf("%064b\n", 15)
	fmt.Printf("%064b\n", bits.ReverseBytes64(15))
}
