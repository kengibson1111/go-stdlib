# crypto/cipher - ctr

This is an example using CTR in order to encrypt and decrypt. The IV is a random value in the first ciphertext block. This example does not cover the use of HMAC after encryption for authentication purposes. The AES encryption algorithm is used. A key distinction for CTR seems to be the ability to be used for both encryption and decryption which is why the sample shows both. For CBC and CFB, encrypters and decrypters are unique.
