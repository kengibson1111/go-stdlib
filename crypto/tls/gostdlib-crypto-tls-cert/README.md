# crypto/tls - cert

This generates a cert and private key - the default host target is localhost, but that can be changed on an input param `--host`. It may help to look at the crypto/rsa and crypto/x509 samples first.

Move server.key and server.pem into the same directory where you will run the TLS server. Take the contents of the cert and copy it into the rootPEM value in the TLS client. Now you have a client that recognizes a cert specific to the TLS server.
