package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", bytes.TrimSpace([]byte(" \t\n a lone gopher \n\t\r\n")))
}
