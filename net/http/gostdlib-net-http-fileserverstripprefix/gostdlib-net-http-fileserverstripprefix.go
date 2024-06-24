package main

import (
	"log"
	"net/http"
)

func main() {
	// To serve a directory on disk (/tmp) under an alternate URL
	// path (/man1/), use StripPrefix to modify the request
	// URL's path before the FileServer sees it:
	http.Handle("/man1/", http.StripPrefix("/man1/", http.FileServer(http.Dir("/usr/share/man"))))
	log.Fatal(http.ListenAndServe(":8180", nil))
}
