# crypto/cipher - gcmencrypter

This is an example using GCM in order to encrypt. A key is used to first create the cipher block, then the cipher block is used to create a GCM encrypter/decrypter. The nonce is a random value used with the GCM encrypter/decrypter to encrypt/decrypt the ciphertext. This example does not cover the use of HMAC after encryption for authentication purposes. The AES encryption algorithm is used.
