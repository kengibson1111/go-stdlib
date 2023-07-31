package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.IndexAny([]byte("go gopher gopher"), "MüQp"))
	fmt.Println(bytes.LastIndexAny([]byte("go gopher gopher"), "MüQp"))

	fmt.Println(bytes.IndexAny([]byte("go 地鼠 地鼠"), "地大"))
	fmt.Println(bytes.LastIndexAny([]byte("go 地鼠 地鼠"), "地大"))

	fmt.Println(bytes.LastIndexAny([]byte("go gopher"), "z,!."))
}
