package main

import (
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"os"
)

func main() {

	cert, err := x509.ParseCertificate(block.Bytes)

	blockType := "RSA PRIVATE KEY"
	password := []byte("password")

	// see http://golang.org/pkg/crypto/x509/#pkg-constants
	cipherType := x509.PEMCipherAES256

	EncryptedPEMBlock, err := x509.EncryptPEMBlock(rand.Reader,
		blockType,
		[]byte("secret message"),
		password,
		cipherType)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// check if encryption is successful or not

	if !x509.IsEncryptedPEMBlock(EncryptedPEMBlock) {
		fmt.Println("PEM Block is not encrypted!")
		os.Exit(1)
	}

	if EncryptedPEMBlock.Type != blockType {
		fmt.Println("Block type is wrong!")
		os.Exit(1)
	}

	fmt.Printf("Encrypted block \n%v\n", EncryptedPEMBlock)

	fmt.Printf("Encrypted Block Headers Info : %v\n", EncryptedPEMBlock.Headers)

	DecryptedPEMBlock, err := x509.DecryptPEMBlock(EncryptedPEMBlock, password)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Decrypted block message is :  \n%s\n", DecryptedPEMBlock)

}
