# go-packages
Code samples using golang's default packages. types.md is a summary of primitive types in golang.
This is a collection of various public examples. If I am working on a project, it is easier to
refer to my own repo and notes.

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
  communication. TO the external world, self-signed certs are evil. But if YOU know where the keys
  and certs are coming from, you can figure out a way to generate keys and certs every 5-10 days.
  And while the external API consumer hates that, internally micro-services are protected through
  a consistent, reliable cert and key rotation strategy. Cool.

## reflect

* makefunc - straight from golang.org's documentation. Still trying to wrap my head around this, but
  I think this may be help with dynamically creating a function and binding to a runtime. Maybe not,
  but an interesting sample.
