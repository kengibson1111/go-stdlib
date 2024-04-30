package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// First number is 33<<32 + 23
	n132 := []uint32{33, 23}
	// Second number is 21<<32 + 12
	n232 := []uint32{21, 12}
	// Sub them together without producing carry.
	d132, carry32 := bits.Sub32(n132[1], n232[1], 0)
	d032, _ := bits.Sub32(n132[0], n232[0], carry32)
	nsum32 := []uint32{d032, d132}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n132, n232, nsum32, carry32)

	// First number is 3<<32 + 2147483647
	n132 = []uint32{3, 0x7fffffff}
	// Second number is 1<<32 + 2147483648
	n232 = []uint32{1, 0x80000000}
	// Sub them together producing carry.
	d132, carry32 = bits.Sub32(n132[1], n232[1], 0)
	d032, _ = bits.Sub32(n132[0], n232[0], carry32)
	nsum32 = []uint32{d032, d132}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n132, n232, nsum32, carry32)

	// First number is 33<<64 + 23
	n164 := []uint64{33, 23}
	// Second number is 21<<64 + 12
	n264 := []uint64{21, 12}
	// Sub them together without producing carry.
	d164, carry64 := bits.Sub64(n164[1], n264[1], 0)
	d064, _ := bits.Sub64(n164[0], n264[0], carry64)
	nsum64 := []uint64{d064, d164}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n164, n264, nsum64, carry64)

	// First number is 3<<64 + 9223372036854775807
	n164 = []uint64{3, 0x7fffffffffffffff}
	// Second number is 1<<64 + 9223372036854775808
	n264 = []uint64{1, 0x8000000000000000}
	// Sub them together producing carry.
	d164, carry64 = bits.Sub64(n164[1], n264[1], 0)
	d064, _ = bits.Sub64(n164[0], n264[0], carry64)
	nsum64 = []uint64{d064, d164}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n164, n264, nsum64, carry64)
}
