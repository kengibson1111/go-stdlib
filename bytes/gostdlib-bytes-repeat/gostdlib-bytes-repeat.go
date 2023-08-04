package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("ba%s\n", bytes.Repeat([]byte("na"), 2))
}
