package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"fmt"
)

func main() {
	key := []byte("example key 1234")
	plaintext := []byte("some plaintext")

	fmt.Printf("Block size: %v\n", aes.BlockSize)

	// should be 32-bytes because the plain text is 16 runes.
	fmt.Printf("Before: %s\n", plaintext)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	fmt.Printf("Decoded: %s\n", ciphertext)

	encodedString := hex.EncodeToString(ciphertext)
	fmt.Printf("Encoded: %s\n", encodedString)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
}
