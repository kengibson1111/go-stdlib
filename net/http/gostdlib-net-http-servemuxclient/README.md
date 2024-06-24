# net/http - servemuxclient

This establishes a TLS handshake and does an HTTP GET on testserver:8443 - the port where `gostdlib-net-http-servemux` is running within a Docker container named `testserver` connected to the same local Docker network. Use the `crypto/tls/gostdlib-crypto-tls-cert` sample to build `server.pem` and `server.key` files.

In my local environment, I run everything in Docker containers and I have a local Docker network so that containers can communicate. After a `go install` in `crypto/tls/gostdlib-crypto-tls-cert`, the executable is available within a container on the container's PATH. In the `gostdlib-net-http-servemuxclient` directory, the command is:

```bash
gostdlib-crypto-tls-cert
```

This generates `server.pem` and `server.key` files. Rename `server.pem` and `server.key` to `client.pem` and `client.key` respectively. Copy the `server.pem` from `gostdlib-net-http-servemux`. This shows how 4 different URLs are handled by `gostdlib-net-http-servemux`. 

## why is this important?

managing 2-way TLS between Docker containers in a Kubernetes pod can be a headache. As part of a pod deployment, the `server.pem` and `server.key` files can be generated for only `127.0.0.1` and the hostname of the container within the pod or the hostname of the pod depending on deployment implementation.

Each "client" micro-service container that needs to communicate with the "server" micro-service container within the pod namespace can have `client.pem` and `client.key` files and the `server.pem` of the "server" container. The client code (based on the sample) has a CA cert pool limited to just the `server.pem`. This is a nice security feature. It also simplifies operations and reduces cost for cert and key management.

To the external world of publicly accessible API endpoints, self-signed certs are evil. However for internal pod communication, if YOU know where the keys and certs are coming from, you can figure out a way to generate keys and certs every 5-10 days. And while the external API consumer hates that, internally micro-services are protected through a consistent, reliable cert and key rotation strategy.
