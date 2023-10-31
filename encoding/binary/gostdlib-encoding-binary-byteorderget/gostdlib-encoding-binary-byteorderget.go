package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	b := []byte{0xe8, 0x03, 0xd0, 0x07}

	x1 := binary.LittleEndian.Uint16(b[0:])
	x2 := binary.LittleEndian.Uint16(b[2:])
	fmt.Printf("little: %#04x %#04x\n", x1, x2)

	x3 := binary.BigEndian.Uint16(b[0:])
	x4 := binary.BigEndian.Uint16(b[2:])
	fmt.Printf("big:    %#04x %#04x\n", x3, x4)
}
