package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	key := []byte("example key 1234")
	encodedtext := "22277966616d9bc47177bd02603d08c9a67d5380d0fe8cf3b44438dff7b9"
	ciphertext, _ := hex.DecodeString(encodedtext)

	fmt.Printf("Block size: %v\n", aes.BlockSize)
	fmt.Printf("Encoded: %s\n", encodedtext)
	fmt.Printf("Decoded: %s\n", ciphertext)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	fmt.Printf("After: %s", ciphertext)
}
