package main

import (
	"archive/zip"
	"os"
	"log"
)

func zipit(target string) {
	zipfile, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}
	defer zipfile.Close()

	w := zip.NewWriter(zipfile)

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
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	zipit("temp.zip")
}
