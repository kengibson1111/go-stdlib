# crypto/tls - client

Generate a server.key and server.pem using **gostdlib-crypto-tls-cert**. Move server.key and server.pem into the same directory where you will run the TLS server. Take the contents of server.pem and copy it into the rootPEM value in the TLS client. Now you have a client that recognizes
a cert specific to tlsserver. The cert and key are only good for localhost so don't try to use the TLS server, and the TLS server will recognize the client.

In the client code, make sure the code formatting is removed from server.pem content. There can't be any spaces or tabs between the lines of certificate content. Start the TLS server. The TLS server will consume the server.key and server.pem files. Run the TLS client. The client will send "hello" to the server and receive "world" from the server.
