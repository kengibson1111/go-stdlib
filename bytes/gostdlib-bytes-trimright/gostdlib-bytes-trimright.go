package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(string(bytes.TrimRight([]byte("453gopher8257"), "0123456789")))
}
