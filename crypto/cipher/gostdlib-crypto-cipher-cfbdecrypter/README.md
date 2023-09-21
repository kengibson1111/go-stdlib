# crypto/cipher - cfbdecrypter

This is an example using CFB to order to decrypt encrypted CFB data. From the output, notice the difference between what is hex-encoded, hex-decoded, and unencrypted. In this case, the first block of ciphertext has the IV for the rest of the ciphertext. This example does not cover the use of HMAC in order to authenticate before decrypting. The AES encryption algorithm is used.
