package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// First, we create a relative symlink to a file.
	d, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(d)

	targetPath := filepath.Join(d, "hello.txt")
	if err := os.WriteFile(targetPath, []byte("Hello, Gophers!"), 0644); err != nil {
		log.Fatal(err)
	}

	linkPath := filepath.Join(d, "hello.link")
	if err := os.Symlink("hello.txt", filepath.Join(d, "hello.link")); err != nil {
		if errors.Is(err, errors.ErrUnsupported) {
			// Allow the example to run on platforms that do not support symbolic links.
			fmt.Printf("%s links to %s\n", filepath.Base(linkPath), "hello.txt")
			return
		}

		log.Fatal(err)
	}

	// Readlink returns the relative path as passed to os.Symlink.
	dst, err := os.Readlink(linkPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s links to %s\n", filepath.Base(linkPath), dst)

	var dstAbs string
	if filepath.IsAbs(dst) {
		dstAbs = dst
	} else {
		// Symlink targets are relative to the directory containing the link.
		dstAbs = filepath.Join(filepath.Dir(linkPath), dst)
	}

	// Check that the target is correct by comparing it with os.Stat
	// on the original target path.
	dstInfo, err := os.Stat(dstAbs)
	if err != nil {
		log.Fatal(err)
	}

	targetInfo, err := os.Stat(targetPath)
	if err != nil {
		log.Fatal(err)
	}

	if !os.SameFile(dstInfo, targetInfo) {
		log.Fatalf("link destination (%s) is not the same file as %s", dstAbs, targetPath)
	}
}
