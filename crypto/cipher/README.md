# crypto/cipher

Terms to understand:

* cipher: method of encrypting data.
* ciphertext: the encrypted data resulting from using a cipher.
* algorithm: a mathematical way to solve a problem.
* key: a value used with an algorithm to encrypt data. Length is important. Think of someone trying
  to break an encryption. They will have to use all key values of all lengths against all possible
  algorithms in order to produce something that is understandable from encrypted data.
* block cipher: a way to encrypt a block of data (like 64 contiguous bits instead of just 1 bit)
  using a cryptographic key and an algorithm.
* stream cipher: basically encrypting one bit at a time. Generally not used, but there are
  use cases. CFB works with streams (see below).
* initialization vector (IV): sometimes called a nonce. It is an arbitrary number chosen in a data
  exchange session. It is used with the key. So to extend the notes of 'key' above, now a hacker
  has to combine all possible IVs with all possible key values of all possible lengths against all
  possible algorithms in order to produce something understandable from encrypted data. IVs are used
  typically during secure browser sessions. SSL and TLS does this behind the scenes using certs
  and keys.
* cipher block chaining (CBC): this is really cool. This is a combination of block cipher and IV but it also
  gets better. The decryption of each block depends on preceding ciphertext blocks - that is the
  chaining. So if there is any bit inconsistency in ANY block for an encrypted message, the entire
  message is considered invalid.
* ciphertext feedback (CFB): like CBC, this is another way to work with block ciphers. But instead
  of encrypting a set number of bits at a time, CFB encrypts one bit at a time. It uses an IV.
  The previously encrypted ciphertext block is XOR'ed with the current plaintext block in order to
  create the current ciphertext block. This works well with streams.
* counter (CTR) - this looks a lot like CFB - one bit at a time, the XOR'ing. The CTR can use an
  IV as a nonce, but also uses an internal counter. And it does not appear that padding to the block
  size is required.
* output feedback (OFB) - this is like CFB, but the difference is that the resulting encrypted block is
  not the ciphertext. It is the output of the XOR. So there are no chaining dependencies. And with no chaining
  dependencies, it is like CTR since OFB can be used for encryption and decryption.
