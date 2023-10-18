package main

import (
	"crypto/tls"
	"log"
)

func main() {
	conn, err := tls.Dial("tcp", "mail.google.com:443", nil)
	if err != nil {
		log.Fatalln("failed to connect: " + err.Error())
	}

	log.Println("connection opened")
	err = conn.Close()
	if err != nil {
		log.Fatalln("failed to close connection: " + err.Error())
	}

	log.Println("connection closed")
}
