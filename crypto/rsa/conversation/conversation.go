package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	// set up the public/private key pairs for the conversation.
	party1PrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	party1PublicKey := &party1PrivateKey.PublicKey

	party2PrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	party2PublicKey := &party2PrivateKey.PublicKey

	fmt.Println("Party1 Private Key : ", party1PrivateKey)
	fmt.Println("Party1 Public key ", party1PublicKey)
	fmt.Println("Party2 Private Key : ", party2PrivateKey)
	fmt.Println("Party2 Public key ", party2PublicKey)

	// now that each party has public/private key pairs, the first issue is
	// how each party gets the public key of the other. With TLS, this is
	// part of the handshake. Research the TLS handshake and it will make
	// sense. This sample is not meant to cover that.

	// with the public key of party2, party1 encrypts a message. Notice the
	// optional label. If party1 wants to send different types of messages
	// with the same public key of party2, the label can be used for switch
	// purposes. It can also be used to authorize the message - if party2
	// cannot see a legal label for a message, it can cancel the entire
	// decryption operation.

	// Also notice that somehow party1 has to let party2 know what the
	// hash algorithm is going to be.
	message := []byte("if you vote for me, all of your dreams will come true")  
	label := []byte("")  
	hash := sha512.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, party2PublicKey, message, label)
	if err != nil {  
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Before encryption [%s]\n", string(message))
	fmt.Printf("After encryption [%x]\n", ciphertext)

	// Now party1 has to sign the message with its private key. party2 is expected
	// to decrypt with party1's public key. Remember the public key exchange? And
	// the random number for the signature does not have to be the same as the
	// random number for the message encryption. Sweet.
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for a basic sample

	newhash := crypto.SHA512
	sigHash := newhash.New()  
	sigHash.Write(message)  
	sum := sigHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, party1PrivateKey, newhash, sum, &opts)
	if err != nil {  
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("PSS signature : %x\n", signature)

	// the "sending" part is done. Just want to reemphasize that party1 and party2 have to
	// be in synch on a few things before even encryption is possible. They each have to
	// have the other's public key. And they each have to agree on a hash algorithm that
	// will be used for encryption. The same algorithm must be used in the signature.

	// you might be wondering, "when I use a browser to hit a TLS site, I don't have a customized
	// public key that I pass to the server as part of the TLS handshake." When using one-way
	// SSL/TLS, the browser must be passing a public key from the O/S root cert chain
	// on the device. Most likely it is coming from a cert for the browser product. At least,
	// it seems logical that this is what's happening in one-way TLS.

	// time for party2 to verify the signature and decrypt.

	// decryption by party2 first. Party2's private key is used. And because
	// party1 and party2 agreed on the hash algorithm and label, it's all cool.
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, party2PrivateKey, ciphertext, label)
	if err != nil {  
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Before decryption [%x]\n", ciphertext)
	fmt.Printf("After decryption [%s]\n", plainText)

	// but the fun is not over yet. Party2 has to verify that the message really came
	// from party1. That is the signature verification using party1's public key.
	err = rsa.VerifyPSS(party1PublicKey, newhash, sum, signature, &opts)
	if err != nil {  
		fmt.Println("Who are f*** are you?! Signature is bogus!!")
		os.Exit(1)
	} else {
		fmt.Println("signature is good")
	}
}
