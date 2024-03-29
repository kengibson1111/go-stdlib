package main

import (
	"archive/zip"
	"errors"
	"log"
	"os"
)

var (
	errNilWriter = errors.New("zip.NewWriter: nil writer")
)

func zipit(target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	w := zip.NewWriter(zipfile)
	if w == nil {
		return errNilWriter
	}

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			w.Close() // ignore Close() error and return the Create() error
			return err
		}

		_, err = f.Write([]byte(file.Body))
		if err != nil {
			w.Close() // ignore Close() error and return the Write() error
			return err
		}
	}

	// Check the error on Close.
	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := zipit("temp.zip")
	if err != nil {
		log.Fatal(err)
	}
}
