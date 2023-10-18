# crypto/tls - dial

The **gostdlib-crypto-tls-cert**, **gostdlib-crypto-tls-server**, and **gostdlib-crypto-tls-client** modules show how to use a custom server.key and server.pem between client and server. This can be useful when managing an "internal only" set of certificates with a very limited scope of use - for example, mTLS connections within a single Kubernetes pod design.

This example shows how to dial a TLS connection with the underlying host's root-CA certificate set.
