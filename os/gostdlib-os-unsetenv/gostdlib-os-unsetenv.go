package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("TMPDIR", "/my/tmp")
	fmt.Println("TMPDIR = " + os.Getenv("TMPDIR"))

	os.Unsetenv("TMPDIR")
	fmt.Println("TMPDIR = " + os.Getenv("TMPDIR"))
}
