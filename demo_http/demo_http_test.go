package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_https(t *testing.T) {
	caCert, err := ioutil.ReadFile("rootCA.crt")
	if err != nil {
		fmt.Println(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	clientCert, err := tls.LoadX509KeyPair("/path/to/client.crt", "/path/to/client.key")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{clientCert},
			},
		},
	}

	_, err = client.Get("https://secure.domain.com")
	if err != nil {
		panic(err)
	}
}
