package main

import (
    "log"
    "crypto/tls"
    "crypto/x509"
)

func main() {
    log.SetFlags(log.Lshortfile)

    rootPEM := `<put server.pem contents here>`

    roots := x509.NewCertPool()
    ok := roots.AppendCertsFromPEM([]byte(rootPEM))
    if !ok {
        panic("failed to parse root certificate")
    }
    tlsConf := &tls.Config{RootCAs: roots}
    conn, err := tls.Dial("tcp", "127.0.0.1:443", tlsConf)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    n, err := conn.Write([]byte("hello\n"))
    if err != nil {
        log.Println(n, err)
        return
    }

    buf := make([]byte, 100)
    n, err = conn.Read(buf)
    if err != nil {
        log.Println(n, err)
        return
    }

    println(string(buf[:n]))
}

