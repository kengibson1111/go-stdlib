# go-stdlib ![go-stdlib](/images/google-go.png)

Code examples using [Go's standard library](https://pkg.go.dev/std).[^1] [Types](https://github.com/kengibson1111/go-stdlib/blob/master/types.md) is a summary of primitive types in Go. This is a collection of various public examples. It is good to know what Go offers out of the box before bulding addition modules. The directories in this repo match the directory structure in the standard library.

Before diving into this repo, it may help to start with the [Tour of Go](https://go.dev/tour/welcome/1). A companion repo for the tour is [here](https://github.com/kengibson1111/tour-of-go).

## bufio

* wordscan - uses a default set of delimiters to scan words in a string.

* writer - shows how to write to a stream.

## bytes

* buffer - shows how to use a buffer in various ways (Write, as a stream, transferring to another
  stream)

* bufferreader - shows how to turn a buffer into an io.Reader. Handy.

* compare - basic comparison and equality.

* comparesearch - shows how to find a byte slice in a sorted set of byte slices.

* trimprefix - prefix trimming and substitution.

* trimsuffix - suffix trimming and substitution.

## compress/flate

* freader - simple reader sample.

* fwriter - simple writer sample.

## compress/gzip

* greader - simple reader sample.

* gwriter - simple writer sample.

## compress/lzw

* lreader - simple reader sample.

* lwriter - simple writer sample.

## compress/zlib

* zreader - simple reader sample.

* zwriter - simple writer sample.

## container/heap

* intheap - simple heap of integers. A heap is great for priority queues.

* pqheap - heap for a custom type.

## container/list

* basic - basic doubly linked list.

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

Samples:

* cbcdecrypter - example using CBC in order to decrypt encrypted CBC data. From the output, notice the difference
  between what is hex-encoded, hex-decoded, and unencrypted. In this case, the first block of ciphertext
  has the IV for the rest of the ciphertext. That doesn't have to always be the case, but it is common.
  So if you want to create a custom encryption/decryption strategy, go against the norm. Put the
  IV somewhere else. This example does not cover the use of HMAC in order to authenticate before
  decrypting. The AES encryption algorithm is used.

* cbcencrypter - example using CBC in order to encrypt. The IV is a random value in the first ciphertext
  block. That doesn't have to always be the case, but it is common. So if you want to create a custom
  encryption/decryption strategy, go against the norm. Put the IV somewhere else. This example does not
  cover the use of HMAC after encryption for authentication purposes. The AES encryption algorithm is used.

* cfbdecrypter - example using CFB to order to decrypt encrypted CFB data. From the output, notice the difference
  between what is hex-encoded, hex-decoded, and unencrypted. In this case, the first block of ciphertext
  has the IV for the rest of the ciphertext. This example does not cover the use of HMAC in order to
  authenticate before decrypting. The AES encryption algorithm is used.

* cfbencrypter - example using CFB in order to encrypt. The IV is a random value in the first ciphertext
  block. This example does not cover the use of HMAC after encryption for authentication purposes. The AES
  encryption algorithm is used.

* ctr - example using CTR in order to encrypt and decrypt. The IV is a random value in the first ciphertext
  block. This example does not cover the use of HMAC after encryption for authentication purposes. The AES
  encryption algorithm is used. A key distinction for CTR seems to be the ability to be used for both
  encryption and decryption which is why the sample shows both. For CBC and CFB, encrypters and decrypters
  are unique.

* ofb - example using OFB in order to encrypt and decrypt. The IV is a random value in the first ciphertext
  block. This example does not cover the use of HMAC after encryption for authentication purposes. The AES
  encryption algorithm is used.

* streamwriter - this shows how to get a plaintext stream and encrypt while writing using OFB. This example does
  not cover the use of HMAC after encryption for authentication purposes. The AES encryption algorithm is used.
  Pretty cool to show off, but not realistic because of lack of authentication.

* streamreader - this shows how to get an encrypted stream and decrypt while writing using OFB. This example does
  not cover the use of HMAC after encryption for authentication purposes. The AES encryption algorithm is used.
  Pretty cool to show off, but not realistic because of lack of authentication.

## crypto/des

* cbcdecrypter - example using CBC in order to decrypt encrypted CBC data. From the output, notice the difference
  between what is hex-encoded, hex-decoded, and unencrypted. In this case, the first block of ciphertext
  has the IV for the rest of the ciphertext. That doesn't have to always be the case, but it is common.
  So if you want to create a custom encryption/decryption strategy, go against the norm. Put the
  IV somewhere else. This example does not cover the use of HMAC in order to authenticate before
  decrypting. Triple DES encryption algorithm is used.

* cbcencrypter - example using CBC in order to encrypt. The IV is a random value in the first ciphertext
  block. That doesn't have to always be the case, but it is common. So if you want to create a custom
  encryption/decryption strategy, go against the norm. Put the IV somewhere else. This example does not
  cover the use of HMAC after encryption for authentication purposes. Triple DES encryption algorithm is used.

## crypto/hmac

This builds upon the hash algorithm samples. Encryption/decryption without authentication allows a
"man in the middle" to modify message bits, recompute a checksum with a hash algorithm, and send both
to a receiver. The hacker may not care what the message contents are - the intent is to wreak havoc.
All of the samples in crypto/cipher do not include authentication.

MAC is what prevents message modification by using a symmetric key between sender and receiver. HMAC is the most common
implementation of MAC because it combines a hash algorithm with a shared key in order to produce
a checksum. The "man in the middle" may be able to eventually guess the hash algorithm producing
a plain checksum, but it is very difficult to predict an HMAC. SSL and TLS use HMAC under the covers
to authenticate encrypted messages after a handshake is established.

It is also very difficult to establish a valid SSL or TLS handshake and have a valid symmetric key
in order to produce an encrypted message with an HMAC. Even so, the symmetric key makes HMAC valid
in limited use cases like SSL or TLS. With evolving technology, it is better to use PKI and
mathematically matched public/private key pairs for custom message encryption/decryption along
with a MAC checksum or digital signature.

The sample does not include encryption/decryption of the message. That would be an enhancement using
samples from crypto/cipher.

## crypto/md5

This shows 2 ways to create the same checksum using a md5 hash.

## crypto/rand

* basic - shows how to fill a byte array with secure random values. Great for initialization vector (IV)
creation.

* uuid - shows how to build a version 4 UUID.

## crypto/rsa

This is pretty cool because it combines functionality from other crypto packages to implement
RSA encryption as specified by PKCS#1 - public/private key pairs which are mathematically
related. Private keys are not distributed and used to decrypt messages and send signatures.
Public keys are distributed and used to encrypt messages and verify signatures.

RSA-OEAP is the latest RSA encryption implementation of the PKCS#1 spec. RSA-PSS is the latest
signature implementation of the PKCS#1 spec. So examples will only focus on those.

Think of the variations possible with OEAP:
(1) It uses an optional label with the message so that if the same public key is used to
encrypt different types of messages, you can verify the label and switch off the label for
decryption.

(2) It uses a random number generator from crypto/rand to make sure 2 encryptions of the
same message with the same hash algorithm are different.

(3) it uses a hash algorithm. The samples use crypto/sha512.

Samples:

* encryptoeap - this shows how to use the public key of an rsa.PrivateKey in order to encrypt.
  A signature is not included.

* decryptoeap - this shows how to use an rsa.PrivateKey in order to decrypt. A signature is
  not included.

* conversation - this shows the basics of both OEAP and PSS. All from public sources.

## crypto/sha1

This shows 2 ways to create the same checksum using a sha1 hash.

## crypto/sha256

This shows 2 ways to create the same checksum using a sha256 hash. Also covers the use of sha224.

## crypto/sha512

This shows 2 ways to create the same checksum using a sha512 hash. Also covers the use of sha384,
sha512/224, and sha512/256.

## crypto/tls
A little prework required for the server and client.

* tlscert - generates a cert and private key for localhost. Straight from golang's public open source.
  I just filled in the blanks and changed isCA to true. Maybe look at the crypto/rsa and crypto/x509
  samples first.

Move server.key and server.pem into the same directory where you will run tlsserver. Take the contents
of the cert and copy it into the rootPEM value in tlsclient. Now you have a client that recognizes
a cert specific to tlsserver. The cert and key are only good for localhost so don't try to use
it externally. It's a demo, man.

* tlsserver - simple example of a TLS server. It requires a generated server.key and server.pem.
  This came from a public gist with few mods.

* tlsclient - simple example of a TLS client. It requires the contents of the server.pem. This
  came from a public gist with a couple of mods to recognize the cert.

## crypto/x509

* verify - basic example showing how to verify a cert against a root cert chain. It also uses
  encoding/pem. This sample will panic because the cert is not valid.

## encoding/base32

* decode - shows how to decode a base32-encoded value to it's original string value.

* encode - shows how to encode a string value to base32.

* newencoder - this streams a string to a base32 value.

## encoding/base64

* package - basic encode/decode sample.

* decode - shows how to decode a base64-encoded value to it's original string value.

* encode - shows how to encode a string value to base64.

* newencoder - this streams a string to a base64 value.

## encoding/binary

* binaryread - shows how to read binary bytes into a float64. That was just for the sample.
  Binary encoding can be used for any type.

* binarywrite - shows how to write a float64 into binary bytes. Reverse of the previous sample.

* binarywritemulti - like the previous example, but shows how to mix various types in a
  single write stream. Printing out the binary bytes actually is a legible string in the sample.

## encoding/csv

* csvread - basic comma-delimited reading.

* csvreadall - like the previous sample only it will read all rows at once.

* csvreadopt - this shows how to set parse and comment options.

* csvwrite - basic writing into comma-delimited rows.

* csvwriteall - like the previous sample only it will write all rows at once.

## encoding/gob

* basic - shows how to use basic functionality. In this sample, you have to imagine one type
  being available in one process for encoding and another version of the same type being
  available in another process. Bytes are streamed and even though the encoded type is different
  from the decoded type, the Decoder gets as much info as is available without an error. Like
  protobufs but it is built into golang.

* custommarshal - shows how to implement gob marshaling for a custom type - necessary because
  internal field types may not have an exported scope.

* custominterface - really cool. The custom type would be registered in the encoder and decoder
  gob processes. The type implements a method on a custom interface. The interface is available
  in the encoder and decoder processes. The encoder encodes through the interface. The decoder
  decodes through the interface and calls the interface's method.

## encoding/json

* decoder - basic JSON decoding

* decoderstream - like the previous example but using the Decoder stream functionality.

* decodertoken - like the previous example but reading individual tokens from the stream.

* indent - shows marshaling and indenting as two separate steps.

* marshal - shows basic marshaling.

* marshalindent - shows marshaling and indenting as one step.

* rawmessage - this is really cool. Shows how to delay parsing until a lower-level type
  is identified. Then parsing continues.

* unmarshal - shows basic unmarshaling.

* unmarshalforward - shows how the JSON encoder is very tolerant. In this case, a component
  was built and deployed with an event type definition. Over time, the event type
  definition evolves, but the component owner has decided to continue using an older type
  definition.
  
  The first JSON blob has a field not included in the current event's type definition. No error on the unmarshal.
  The second JSON blob is the next version of the first JSON blob - 2 versions ahead
  of the current event's type definition. In this case, a field is no longer used. Again,
  no error on the unmarshal. If the missing field is required by the component, it
  can return an error related to the missing data. And the component owner can decide
  how to handle the three different versions of JSON blobs.

* unmarshalbackward - also shows JSON encoder tolerance. In this case, components are still
  publishing version 1 events. But there are subscribers with version 2 and version 3 event
  type definitions. No error on the unmarshal, and the subscriber component's owner can decide
  if/when to return an application error for previous versions of events.

## encoding/xml

* encoder - this shows how to define xml translation in a golang type definition. That information
  is used by an Encoder to output XML to a stream.

* marshalindent - MarshalIndent works like Marshal, but prettifies the output byte slice.

* unmarshal - like the previous 2 samples with type definition, but reading XML instead of
  writing.

* soapdecode - this shows how to decode a simple SOAP envelope with a header and body. The
  XSI and XSD namespaces were included as attrs in the envelope, but never used. Just trying
  to show that the golang parser isn't going to choke on an unused namespace. At some point,
  you need to be able to get to individual tokens in a header or body and then apply validation
  rules against those tokens. This sample gets to the individual tokens, but does no validation.
  That would be an extension to the sample.

## errors

* basic - basic example of a custom implementation of the error interface. tour-of-go covers
  interface implementations (including error) in more detail.

* newerror - the only sample which actually imports the errors package. Still not sure
  how useful this is compared to the basic sample.

* errorf - this uses fmt.Errorf to create a formatted error.

##flag

* basic - this shows how to pass command-line args utilizing the flag package. It shows how to
  define a flag through a pointer, how to bind a variable to a flag definition instead of
  the pointer option, and how to bind a variable to multiple flag entries. Pretty cool.

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

* printints - shows several formatting verbs for integers.

* printfloats - shows several formatting verbs for floats.

* printptr - just demonstrates %p.

## html

* escapestring - special escaping for only <, >, &, ', and ".

* unescapestring - special unescaping for only <, >, &, ', and ".

## html/template

* basic - shows basic template functionality.

* block - this is showing how to customize a template block. block is shorthand for defining a
  template and executing it in place. That is what allows the template block to be redefined.
  In this sample, overlay is redefining the master template block. Notice how the master template
  instance assign the join function pointer and not the overlay. But the overlay template
  redefinition of the master template block uses join. Starting to get interesting when
  considering dynamic execution functionality. This is the same as the text/template block
  sample.

* autoescape - this shows some of the safety features of html/template vs. text/template.
  HTML escaping happens by default. So it looks like it combines functionality from the
  default html package with text/template.

* escape - this shows HTML, JS, and URL escaping.

## log

* logger - basic use of a Logger.

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

***

[^1]: Icon provided by [seekvectors.com](https://seekvectors.com/).
