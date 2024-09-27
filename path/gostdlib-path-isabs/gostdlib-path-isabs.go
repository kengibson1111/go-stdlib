package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.IsAbs("/dev/null"))
	fmt.Println(path.IsAbs("./temp"))
}
