package main

import (
	"archive/tar"
	"bytes"
	"io"
	"log"
	"os"
)

func main() {
	// Create and add some files to the archive.
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	if tw == nil {
		log.Fatal("tar.NewWriter: nil writer")
	}

	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal("tw.WriteHeader: ", err)
		}

		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal("tw.Write: ", err)
		}
	}

	if err := tw.Close(); err != nil {
		log.Fatal("tw.Close: ", err)
	}

	// Open and iterate through the files in the archive.
	tr := tar.NewReader(&buf)
	if tr == nil {
		log.Fatal("tar.NewReader: nil reader")
	}

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive, gracefully handle error
		}

		if err != nil {
			log.Fatal("tr.Next: ", err)
		}

		log.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal("io.Copy: ", err)
		}

		log.Println()
	}
}
