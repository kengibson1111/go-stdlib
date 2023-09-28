# crypto/des - cbcencrypter

This is an example using CBC in order to encrypt. The IV is a random value in the first ciphertext block. That doesn't have to always be the case, but it is common. So if you want to create a custom encryption/decryption strategy, go against the norm. Put the IV somewhere else. This example does not cover the use of HMAC after encryption for authentication purposes. Triple DES encryption algorithm is used.
