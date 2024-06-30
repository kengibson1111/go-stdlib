package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, TLS!\n")
	})

	log.Printf("About to listen on 8443.")
	err := http.ListenAndServeTLS(":8443", "server.pem", "server.key", nil)
	log.Fatal(err)
}
