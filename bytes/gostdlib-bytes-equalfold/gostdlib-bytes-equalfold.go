package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go")))
	fmt.Println(bytes.Equal(nil, []byte("")))
}
