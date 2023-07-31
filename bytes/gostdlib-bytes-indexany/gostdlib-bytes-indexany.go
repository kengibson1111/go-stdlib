package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.IndexAny([]byte("go 地鼠 地鼠"), "地大"))
	fmt.Println(bytes.IndexAny([]byte("crwth"), "aeiouy"))
}
