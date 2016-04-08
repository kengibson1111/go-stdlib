package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// Load client cert
	cert, err := tls.LoadX509KeyPair("client.pem", "client.key")
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile("server.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// try the home path.
	resp, err := client.Get("https://127.0.0.1:443/")
	if err != nil {
		fmt.Println(err)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", string(contents))

	// try an invalid path
	resp, err = client.Get("https://127.0.0.1:443/hello")
	if err != nil {
		fmt.Println(err)
	}
	contents, err = ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", string(contents))

	// try valid paths
	resp, err = client.Get("https://127.0.0.1:443/api/v1/someapi")
	if err != nil {
		fmt.Println(err)
	}
	contents, err = ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", string(contents))

	resp, err = client.Get("https://127.0.0.1:443/api/v1/someapi2")
	if err != nil {
		fmt.Println(err)
	}
	contents, err = ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", string(contents))
}