package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, " And Leon's getting laaarger!")
	fmt.Printf("%x\n", h.Sum(nil))

	data := []byte("The fog is getting thicker! And Leon's getting laaarger!")
	fmt.Printf("%x\n", md5.Sum(data))
}
