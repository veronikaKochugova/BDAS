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

	// // see http://golang.org/pkg/crypto/x509/#pkg-constants
	// cipherType := x509.PEMCipherAES256

	// EncryptedPEMBlock, err := x509.EncryptPEMBlock(rand.Reader,
	// 	blockType,
	// 	[]byte("secret message"),
	// 	password,
	// 	cipherType)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// // check if encryption is successful or not

	// if !x509.IsEncryptedPEMBlock(EncryptedPEMBlock) {
	// 	fmt.Println("PEM Block is not encrypted!")
	// 	os.Exit(1)
	// }

	// if EncryptedPEMBlock.Type != blockType {
	// 	fmt.Println("Block type is wrong!")
	// 	os.Exit(1)
	// }

	// fmt.Printf("Encrypted block \n%v\n", EncryptedPEMBlock)

	// fmt.Printf("Encrypted Block Headers Info : %v\n", EncryptedPEMBlock.Headers)

	// DecryptedPEMBlock, err := x509.DecryptPEMBlock(EncryptedPEMBlock, password)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("Decrypted block message is :  \n%s\n", DecryptedPEMBlock)

}
