package main

import (
	"crypto/sha256"
	"crypto/hmac"
	"fmt"
)

// createMAC creates a MAC based on the input message and key. This is used by the message sender.
func createMAC(message, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}

// checkMAC reports whether messageMAC is a valid HMAC tag for message. This is used by the message
// receiver
func checkMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func main() {
	// this is a symmetric key shared by sender and receiver of the message.
	key := []byte("example key 1234")

	message := []byte("His money is twice tainted: 'taint yours and 'taint mine.")

	// the sender creates the mac. In this case, an HMAC with a sha256 hash algorithm
	mac := createMAC(message, key)

	// the receiver checks the mac and message from the sender using the shared
	// symmetric key. This is what TLS does behind the scenes after a handshake is established.
	fmt.Printf("MACs match = %v", checkMAC(message, mac, key))
	fmt.Println()
}
