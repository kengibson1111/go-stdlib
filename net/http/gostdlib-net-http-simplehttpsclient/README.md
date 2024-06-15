# net/http - simplehttpsclient

This establishes a TLS handshake and does an HTTP GET on 127.0.0.1:8443 - the port where simplehttpsserver is running. Use the crypto/tls/gostdlib-crypto-tls-cert sample to build server.pem and server.key files. Rename those to client.pem and client.key respectively. Copy the server.pem from simplehttpsserver.
