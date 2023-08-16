package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("[%q]\n", bytes.Trim([]byte(" !!! Achtung! Achtung! !!! "), "! "))
}
