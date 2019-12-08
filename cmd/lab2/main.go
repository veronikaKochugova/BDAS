package main

import (
	"fmt"
	"os"

	lib "./crypto"
	"github.com/akamensky/argparse"
)

func parseArgs() (secret string) {
	parser := argparse.NewParser("main.go", "Encrypt/Decrypt msg")
	s := parser.String("s", "secret", &argparse.Options{Required: true, Help: "secret msg"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *s == "" {
		*s = "Secret msg"
	}
	return *s
}

func main() {

	secret := parseArgs()

	cert := lib.GetCertificate("examples/public.cer")
	keyStore := lib.GetKeyStore("examples/private.p12", "password")

	pubKey := lib.GetPublicKeyFromCert(cert)
	fmt.Println(*pubKey)
	privKey := lib.GetPrivateKeyFromKeyStore(keyStore)

	encrypted := lib.EncryptData(secret, pubKey)
	decrypted := lib.DecryptData(encrypted, privKey)

	fmt.Println("\nSecret msg: " + secret)
	fmt.Println("\nEncrypted msg: " + encrypted)
	fmt.Println("\nDecrypted msg: " + decrypted)
}
