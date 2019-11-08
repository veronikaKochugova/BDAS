package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/pkcs12"
)

func main() {

	// Load public certificate
	buf, err := ioutil.ReadFile("examples/public.cer")
	fmt.Println(err)
	fmt.Println(buf)

	// Decode certificate fron buffer and create block
	block, _ := pem.Decode([]byte(buf))
	if block == nil {
		panic("failed to parse certificate PEM")
	}

	// Parse block and create x509.Certificate instance
	cert, err := x509.ParseCertificate(block.Bytes)

	fmt.Println(err)
	fmt.Println(cert)

	// blockType := "RSA PRIVATE KEY"
	keyPassword := []byte("password")
	// keystorePassword := []byte("password")

	p12, _ := ioutil.ReadFile("examples/private.p12")

	blocks, err := pkcs12.ToPEM(p12, string(keyPassword))
	if err != nil {
		panic(err)
	}
	fmt.Println(blocks)

	// var pemData []byte
	// for _, b := range blocks {
	// 	pemData = append(pemData, pem.EncodeToMemory(b)...)
	// }

	// // then use PEM data for tls to construct tls certificate:
	// cert, err := tls.X509KeyPair(pemData, pemData)
	// if err != nil {
	// 	panic(err)
	// }

	// config := &tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// }

	// _ = config

}
