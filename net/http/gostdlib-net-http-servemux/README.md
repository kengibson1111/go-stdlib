# net/http - servemux

This builds on `gostdlib-net-http-simplehttpsserver`. It uses TLS, so you will need to generate `server.pem` and `server.key` files using the `crypto/tls/gostdlib-crypto-tls-cert` sample. Drop those 2 files in the same directory where you run `gostdlib-net-http-servemux`. You can hit the server using `gostdlib-net-http-servemuxclient`.

```bash
gostdlib-crypto-tls-cert --host "127.0.0.1,testserver"
```

This generates `server.pem` and `server.key` files. I start a container connected to my local Docker network named `testserver` with port 8443 mapped. Then within the container I start `gostdlib-net-http-servemux` in the module's directory (where `server.pem` and `server.key` are located).

This is demonstrating the ability to assign handlers to URL patterns. Patterns are matched from the most specific to least specific. And the advantage over just a function is the ability to define a complex type that implements the Handler interface. In the sample code, the Handler implementation is an empty struct. But it doesn't have to be.
