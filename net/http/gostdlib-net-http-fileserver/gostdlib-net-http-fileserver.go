package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/usr/share/doc")))
	log.Fatal(http.ListenAndServe(":8180", nil))
}
