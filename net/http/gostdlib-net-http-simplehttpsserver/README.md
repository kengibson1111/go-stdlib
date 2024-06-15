# net/http - simplehttpsserver

This listens on 8443 and echos the URL without the leading forward slash. It uses TLS, so you will need to generate server.pem and server.key files using the crypto/tls/gostdlib-crypto-tls-cert sample. Drop those 2 files in the same directory where you run simplehttpsserver. You can hit the server using simplehttpsclient.
