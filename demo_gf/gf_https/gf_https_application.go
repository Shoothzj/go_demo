package main

import (
	"crypto/tls"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
)

func main() {
	server := g.Server()
	server.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/swallow", swallow)
	})
	config, err := LoadCertByPath("gf_https/server.crt", "gf_https/server.key", "pwdpwdpwd")
	if err != nil {
		panic(err)
	}
	server.SetTLSConfig(config)
	server.Run()
}

func LoadCertByPath(certPath, keyPath, pwd string) (*tls.Config, error) {
	certBytes, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, err
	}
	keyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		MinVersion: tls.VersionTLS13,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}, nil
}

func swallow(r *ghttp.Request) {
}
