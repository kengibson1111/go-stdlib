# net/http - simplehttpsserver

This listens on 8443 and echos the URL without the leading forward slash. It uses TLS, so you will need to generate `server.pem` and `server.key` files using the `crypto/tls/gostdlib-crypto-tls-cert` sample. Drop those 2 files in the same directory where you run `gostdlib-net-http-simplehttpsserver`. You can hit the server using `gostdlib-net-http-simplehttpsclient`.

In my local environment, I run everything in Docker containers and I have a local Docker network so that containers can communicate. After a `go install` in `crypto/tls/gostdlib-crypto-tls-cert`, the executable is available within a container on the container's PATH. In the `gostdlib-net-http-simplehttpsserver` directory, the command is:

```bash
gostdlib-crypto-tls-cert --host "127.0.0.1,testserver"
```

This generates `server.pem` and `server.key` files. I start a container connected to my local Docker network named `testserver` with port 8443 mapped. Then within the container I start `gostdlib-net-http-simplehttpsserver` in the module's directory (where `server.pem` and `server.key` are located).
