package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/pkcs12"
)

// SIGN: https://stackoverflow.com/questions/52775864/how-to-read-pkcs12-content-in-golang-i-have-example-in-php
// Verify: https://medium.com/@raul_11817/golang-cryptography-rsa-asymmetric-algorithm-e91363a2f7b3

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}

func GetCertificate(file string) x509.Certificate {
	// Load public certificate to buffer
	buf, err := ioutil.ReadFile(file)
	checkError(err)

	// Decode certificate from buffer and create block
	block, _ := pem.Decode([]byte(buf))

	// Parse block and create x509.Certificate instance
	cert, err := x509.ParseCertificate(block.Bytes)
	checkError(err)

	return *cert
}

func GetKeyStore(file, keystorePassword string) interface{} {

	// Load PKCS12 KeyStore
	buf, err := ioutil.ReadFile(file)
	checkError(err)

	// Get KeyStore.
	keyStore, _, err := pkcs12.Decode(buf, keystorePassword)
	checkError(err)

	return keyStore
}

func GetPrivateKeyFromKeyStore(keyStore interface{}) *rsa.PrivateKey {
	return keyStore.(*rsa.PrivateKey)
}

func GetPublicKeyFromCert(cert x509.Certificate) *rsa.PublicKey {
	return cert.PublicKey.(*rsa.PublicKey)
}

func EncryptData(data string, pubKey *rsa.PublicKey) string {
	dataBytes := []byte(data)
	hash := sha256.New()
	encryptedBytes, err := rsa.EncryptOAEP(hash, rand.Reader, pubKey, dataBytes, nil)
	checkError(err)
	return string(encryptedBytes)
}

func DecryptData(data string, privKey *rsa.PrivateKey) string {
	dataBytes := []byte(data)
	hash := sha256.New()
	decryptedBytes, err := rsa.DecryptOAEP(hash, rand.Reader, privKey, dataBytes, nil)
	checkError(err)
	return string(decryptedBytes)
}
