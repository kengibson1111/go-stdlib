package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%08b\n", 15)
	fmt.Printf("%08b\n", bits.RotateLeft8(15, 2))
	fmt.Printf("%08b\n", bits.RotateLeft8(15, -2))

	fmt.Printf("%016b\n", 15)
	fmt.Printf("%016b\n", bits.RotateLeft16(15, 2))
	fmt.Printf("%016b\n", bits.RotateLeft16(15, -2))

	fmt.Printf("%032b\n", 15)
	fmt.Printf("%032b\n", bits.RotateLeft32(15, 2))
	fmt.Printf("%032b\n", bits.RotateLeft32(15, -2))

	fmt.Printf("%064b\n", 15)
	fmt.Printf("%064b\n", bits.RotateLeft64(15, 2))
	fmt.Printf("%064b\n", bits.RotateLeft64(15, -2))
}
