package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// First number is 0<<32 + 6
	n132 := []uint32{0, 6}
	// Second number is 0<<32 + 3
	n232 := []uint32{0, 3}
	// Divide them together.
	quo32, rem32 := bits.Div32(n132[0], n132[1], n232[1])
	nsum32 := []uint32{quo32, rem32}
	fmt.Printf("[%v %v] / %v = %v\n", n132[0], n132[1], n232[1], nsum32)

	// First number is 2<<32 + 2147483648
	n132 = []uint32{2, 0x80000000}
	// Second number is 0<<32 + 2147483648
	n232 = []uint32{0, 0x80000000}
	// Divide them together.
	quo32, rem32 = bits.Div32(n132[0], n132[1], n232[1])
	nsum32 = []uint32{quo32, rem32}
	fmt.Printf("[%v %v] / %v = %v\n", n132[0], n132[1], n232[1], nsum32)

	// First number is 0<<64 + 6
	n164 := []uint64{0, 6}
	// Second number is 0<<64 + 3
	n264 := []uint64{0, 3}
	// Divide them together.
	quo64, rem64 := bits.Div64(n164[0], n164[1], n264[1])
	nsum64 := []uint64{quo64, rem64}
	fmt.Printf("[%v %v] / %v = %v\n", n164[0], n164[1], n264[1], nsum64)

	// First number is 2<<64 + 9223372036854775808
	n164 = []uint64{2, 0x8000000000000000}
	// Second number is 0<<64 + 9223372036854775808
	n264 = []uint64{0, 0x8000000000000000}
	// Divide them together.
	quo64, rem64 = bits.Div64(n164[0], n164[1], n264[1])
	nsum64 = []uint64{quo64, rem64}
	fmt.Printf("[%v %v] / %v = %v\n", n164[0], n164[1], n264[1], nsum64)
}
