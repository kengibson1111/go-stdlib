package main

import (
	"crypto/ed25519"
	"log"
)

type zeroReader struct{}

func (zeroReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 0
	}

	return len(buf), nil
}

func main() {
	var zero zeroReader
	public, private, err := ed25519.GenerateKey(zero)
	if err != nil {
		log.Fatalln(err)
	}

	// test sign/verify default options
	message := []byte("test message")
	sig := ed25519.Sign(private, message)
	log.Println("message =", message)
	log.Println("sig     =", sig)

	if !ed25519.Verify(public, message, sig) {
		log.Fatalln("valid signature rejected")
	}

	log.Println("message signed and verified...")

	wrongMessage := []byte("wrong message")
	log.Println("wrong message =", wrongMessage)

	if ed25519.Verify(public, wrongMessage, sig) {
		log.Fatalln("signature of different message accepted")
	}

	log.Println("wrong message rejected...")
}
