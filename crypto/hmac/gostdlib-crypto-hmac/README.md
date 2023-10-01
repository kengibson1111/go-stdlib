# crypto/hmac

This builds upon the hash algorithm samples. Encryption/decryption without authentication allows "man in the middle" to modify message bits, recompute a checksum with a hash algorithm, and send both to a receiver. The hacker may not care what the message contents are - the intent is to wreak havoc. All of the samples in crypto/cipher do not include authentication.

MAC is what prevents message modification by using a symmetric key between sender and receiver. HMAC is the most common implementation of MAC because it combines a hash algorithm with a shared key in order to produce a checksum. The "man in the middle" may be able to eventually guess the hash algorithm producing a plain checksum, but it is very difficult to predict an HMAC. SSL and TLS use HMAC under the covers to authenticate encrypted messages after a handshake is established.

It is also very difficult to establish a valid SSL or TLS handshake and have a valid symmetric key in order to produce an encrypted message with an HMAC. Even so, the symmetric key makes HMAC valid in limited use cases like SSL or TLS. With evolving technology, it is better to use PKI and mathematically matched public/private key pairs for custom message encryption/decryption along with a MAC checksum or digital signature.

The sample does not include encryption/decryption of the message. That would be an enhancement using samples from crypto/cipher.
