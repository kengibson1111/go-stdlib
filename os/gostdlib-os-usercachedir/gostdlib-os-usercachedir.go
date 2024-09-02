package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	dir, dirErr := os.UserCacheDir()
	if dirErr == nil {
		dir = filepath.Join(dir, "ExampleUserCacheDir")
	}

	getCache := func(name string) ([]byte, error) {
		if dirErr != nil {
			return nil, &os.PathError{Op: "getCache", Path: name, Err: os.ErrNotExist}
		}

		return os.ReadFile(filepath.Join(dir, name))
	}

	var mkdirOnce sync.Once
	putCache := func(name string, b []byte) error {
		if dirErr != nil {
			return &os.PathError{Op: "putCache", Path: name, Err: dirErr}
		}

		mkdirOnce.Do(func() {
			if err := os.MkdirAll(dir, 0700); err != nil {
				log.Printf("can't create user cache dir: %v", err)
			}
		})

		return os.WriteFile(filepath.Join(dir, name), b, 0600)
	}

	// Read and store cached data.
	_ = getCache
	_ = putCache
}
