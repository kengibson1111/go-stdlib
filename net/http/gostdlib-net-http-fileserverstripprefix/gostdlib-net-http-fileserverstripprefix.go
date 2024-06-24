package main

import (
	"log"
	"net/http"
)

func main() {
	// To serve a directory on disk (/usr/share/man) under an alternate URL
	// path (/manpages/), use StripPrefix to modify the request
	// URL's path before the FileServer sees it:
	http.Handle("/manpages/", http.StripPrefix("/manpages/", http.FileServer(http.Dir("/usr/share/man"))))
	log.Fatal(http.ListenAndServe(":8180", nil))
}
