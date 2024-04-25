package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// First number is 0<<32 + 12
	n132 := []uint32{0, 12}
	// Second number is 0<<32 + 12
	n232 := []uint32{0, 12}
	// Multiply them together without producing overflow.
	hi32, lo32 := bits.Mul32(n132[1], n232[1])
	nsum32 := []uint32{hi32, lo32}
	fmt.Printf("%v * %v = %v\n", n132[1], n232[1], nsum32)

	// First number is 0<<32 + 2147483648
	n132 = []uint32{0, 0x80000000}
	// Second number is 0<<32 + 2
	n232 = []uint32{0, 2}
	// Multiply them together producing overflow.
	hi32, lo32 = bits.Mul32(n132[1], n232[1])
	nsum32 = []uint32{hi32, lo32}
	fmt.Printf("%v * %v = %v\n", n132[1], n232[1], nsum32)

	// First number is 0<<64 + 12
	n164 := []uint64{0, 12}
	// Second number is 0<<64 + 12
	n264 := []uint64{0, 12}
	// Multiply them together without producing overflow.
	hi64, lo64 := bits.Mul64(n164[1], n264[1])
	nsum64 := []uint64{hi64, lo64}
	fmt.Printf("%v * %v = %v\n", n164[1], n264[1], nsum64)

	// First number is 0<<64 + 9223372036854775808
	n164 = []uint64{0, 0x8000000000000000}
	// Second number is 0<<64 + 2
	n264 = []uint64{0, 2}
	// Multiply them together producing overflow.
	hi64, lo64 = bits.Mul64(n164[1], n264[1])
	nsum64 = []uint64{hi64, lo64}
	fmt.Printf("%v * %v = %v\n", n164[1], n264[1], nsum64)
}
