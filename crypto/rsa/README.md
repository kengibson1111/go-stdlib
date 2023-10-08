# crypto/rsa

This is pretty cool because it combines functionality from other crypto packages to implemen RSA encryption as specified by PKCS#1 v2 - public/private key pairs which are mathematically related. Private keys are not distributed and used to decrypt messages and send signatures. Public keys are distributed and used to encrypt messages and verify signatures.

RSA-OEAP is the latest RSA encryption implementation of the PKCS#1 v2 spec. RSA-PSS is the latest signature implementation of the PKCS#1 v2 spec. So examples will only focus on those.

Think of the variations possible with OEAP:
(1) It uses an optional label with the message so that if the same public key is used to encrypt different types of messages, you can verify the label and switch off the label for decryption.

(2) It uses a random number generator from crypto/rand to make sure 2 encryptions of the same message with the same hash algorithm are different.

(3) it uses a hash algorithm. The samples use crypto/sha512.
