# go-packages
Code samples using golang's default packages. types.md is a summary of primitive types in golang.
This is a collection of various public examples. If I am working on a project, it is easier to
refer to my own repo and notes.

## crypto/cipher

Terms to understand:

* cipher: method of encrypting data.
* ciphertext: the encrypted data resulting from using a cipher.
* algorithm: a mathematical way to solve a problem.
* key: a value used with an algorithm to encrypt data. Length is important. Think of someone trying
  to break an encryption. They will have to use all key values of all lengths against all possible
  algorithms in order to produce something that is understandable from encrypted data.
* block cipher: a way to encrypt a block of data (like 64 contiguous bits instead of just 1 bit)
  using a cryptographic key and an algorithm.
* stream cipher: basically encrypting one but at a time. Generally not used, but there are
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

Samples:

* cbcdecrypter - example using CBC in order to decrypt encrypted CBC data. From the output, notice the difference
  between what is hex-encoded, hex-decoded, and unencrypted. In this case, the first block of ciphertext
  has the IV for the rest of the ciphertext. That doesn't have to always be the case, but it is common.
  So if you want to create a custom encryption/decryption strategy, go against the norm. Put the
  IV somewhere else. This example does not cover the use of HMAC in order to authenticate before
  decrypting. The AES encryption algorithm is used.

* cbcencrypter - example using CBC in order to encrypt. The IV is a random value in the first CBC
  block. That doesn't have to always be the case, but it is common. So if you want to create a custom
  encryption/decryption strategy, go against the norm. Put the IV somewhere else. This example does not
  cover the use of HMAC after encryption for authentication purposes. The AES encryption algorithm is used.

* cfbdecrypter - example using CFB to order to decrypt encrypted CFB data. From the output, notice the difference
  between what is hex-encoded, hex-decoded, and unencrypted. In this case, the first block of ciphertext
  has the IV for the rest of the ciphertext. This example does not cover the use of HMAC in order to
  authenticate before decrypting. The AES encryption algorithm is used.

* cfbencrypter - example using CFB in order to encrypt. The IV is a random value in the first CFB
  block. This example does not cover the use of HMAC after encryption for authentication purposes. The AES
  encryption algorithm is used.

## crypto/rand

This shows how to fill a byte array with secure random values. Great for initialization vector (IV)
creation.

## crypto/tls
A little prework required for the server and client.

* tlscert - generates a cert and private key for localhost. Straight from golang's public open source.
  I just filled in the blanks and changed isCA to true.

Move server.key and server.pem into the same directory where you will run tlsserver. Take the contents
of the cert and copy it into the rootPEM value in tlsclient. Now you have a client that recognizes
a cert specific to tlsserver. The cert and key are only good for localhost so don't try to use
it externally. It's a demo, man.

* tlsserver - simple example of a TLS server. It requires a generated server.key and server.pem.
  This came from a public gist with few mods.

* tlsclient - simple example of a TLS client. It requires the contents of the server.pem. This
  came from a public gist with a couple of mods to recognize the cert.

## fmt

Formatted I/O

* printstrings - several samples showing formatting verbs for strings.

* printbytes - same as printstrings only for a byte array.

* rawbytes - more printing but starting with a raw string. Not sure why I would want to use raw strings
  if I have a UTF-8 editor.

* runes - notice the term "char" is not used that much in golang. That is because the Unicode standard
  uses the term "code point". Those include Latin characters we are familiar with and much more. So a
  Unicode code point is a golang "rune". And most of the time it doesn't matter because we usually only
  are interested in Latin characters that can be contained within one byte. We use a for loop one byte
  at a time and everything is fine. We use a for range loop and still everything is fine with Latin
  characters. But with runes, a for range loop will start each iteration at the beginning of the rune
  while a for loop will iterate one byte at a time. Something to think about when processing strings.
  The sample code uses a for range loop. Notice the byte position.

* runesutf8pkg - same results as runes but the code uses the unicode/utf8 package.

* printstructs - a couple of different struct types. %T, %v, %+v seem to be the most effective printing
  verbs for structs.

* printbool - pretty boring :). just demonstrating %t.

## html

* escapestring - special escaping for only <, >, &, ', and ".

* unescapestring - special unescaping for only <, >, &, ', and ".

## net/http

HTTP client and server functions.

* simplehttpserver - this listens on 8080 and echos the URL without the leading forward slash. Straight
  from golang.org's wiki.

* simplehttpsserver - this listens on 443 and echos the URL without the leading forward slash. It uses
  TLS, so you will need to generate server.pem and server.key files using the crypto/tls/tlscert sample.
  Drop those 2 files in the same directory where you run simplehttpsserver. You can hit the server
  using simplehttpsclient.

* simplehttpsclient - this establishes a TLS handshake and does an HTTP GET on 127.0.0.1:443 - the
  port where simplehttpsserver is running. Use the crypto/tls/tlscert sample to build server.pem
  and server.key files. Rename those to client.pem and client.key respectively. Copy the server.pem
  from simplehttpsserver.

* servemux - this builds on simplehttpsserver. It uses TLS, so you will need to generate server.pem and
  server.key files using the crypto/tls/tlscert sample. Drop those 2 files in the same directory where you
  run servemux. You can hit the server using servemuxclient. This is demonstrating the ability to
  assign handlers to URL patterns. Patterns are matched from the most specific to least specific. And
  the advantage over just a function is the abillity to define a complex type that implements the
  Handler interface. In the sample code, the Handler implementation is an empty struct. But
  it doesn't have to be. Very cool.

* servemuxclient - this establishes a TLS handshake and does an HTTP GET on 127.0.0.1:443 - the
  port where servemux is running. Use the crypto/tls/tlscert sample to build server.pem
  and server.key files. Rename those to client.pem and client.key respectively. Copy the server.pem
  from servemux. This shows how 4 different URLs are handled by servemux. What I like about this
  is how you can use self-signed certs and self-generated keys to control internal micro-service
  communication. To the external world, self-signed certs are evil. But if YOU know where the keys
  and certs are coming from, you can figure out a way to generate keys and certs every 5-10 days.
  And while the external API consumer hates that, internally micro-services are protected through
  a consistent, reliable cert and key rotation strategy. Cool.

## reflect

* makefunc - straight from golang.org's documentation. Still trying to wrap my head around this, but
  I think this may help with dynamically creating a function and binding to a runtime. Maybe not,
  but an interesting sample.

* structtag - from golang's docs. It looks like golang provides an easy way to parse JSON-like tags and
  values within a string. Very nice. You can reflect an internal type that is a struct with one string.
  The string value is intended to be like JSON tags and values. Get the one field value through
  reflection. Then get the tags. Maybe a way to extend this example is to look up tag names
  dynamically and grab values.

* typeof - from golang's docs. This shows the use of TypeOf in order to check whether or not an
  interface implementation really implements an interface. Pretty simple.

## testing

* benchparallel - this doesn't output anything at this point, but I thought it was interesting
  because of the text/template work I was doing at the time. Benchmark functions are meant to be run
  as part of a test suite usually with the "go test -cpu" command. golang's baseline text/template
  functionality is very powerful, and I saw this as a value extension to any templates created and
  executed. I will most likely revisit this.

## text/scanner

This shows how to use scanner to parse text tokens.

## text/tabwriter

This shows how to use tabwriter to format text in neat columns.

## text/template

pre-cursor to html/template examples. Everything in html/template builds upon text/template.

* template - shows basic template functionality.

* block - this is showing how to customize a template block. block is shorthand for defining a
  template and executing it in place. That is what allows the template block to be redefined.
  In this sample, overlay is redefining the master template block. Notice how the master template
  instance assign the join function pointer and not the overlay. But the overlay template
  redefinition of the master template block uses join. Starting to get interesting when
  considering dynamic execution functionality.

* func - like the previous sample, this assigns a function pointer. I like this sample because
  it also shows how to use a pipe ("|") feeding the return value of one function into another
  like a Unix pipe.

* glob - this sample shows how to load all templates into a template instance based on a global
  pattern. Then if templates call other templates, they're all available in the same template
  instance. golang has some serious dynamic execution muscle.

* helpers - this builds on the previous sample using ParseGlob to load helper templates. Then
  it dynamically loads two more templates to the helpers and executes each of the new templates.
  Although this sample doesn't show it, imagine if your helpers are calling functions like
  in the block and func samples. Hmmmmm.

* share - this shows how to define a set of base helpers and share them. In this sample, T0
  and T1 are defined and we know that T1 will be calling T2, but T2 is not defined yet. Then
  the helpers are cloned twice - each with its own definition of T2. Kind of sounds like a
  dynamic way to define 2 implementations of the same interface. Powerful.
