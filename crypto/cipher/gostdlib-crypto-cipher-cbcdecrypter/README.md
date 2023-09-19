# crypto/cipher - cbcdecrypter

This is an example using CBC in order to decrypt encrypted CBC data. From the output, notice the difference between what is hex-encoded, hex-decoded, and unencrypted. In this case, the first block of ciphertext has the IV for the rest of the ciphertext. That doesn't have to always be the case, but it is common. So if you want to create a custom encryption/decryption strategy, go against the norm. Put the IV somewhere else. This example does not cover the use of HMAC in order to authenticate before decrypting. The AES encryption algorithm is used.
