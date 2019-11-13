package main

import (
	"fmt"

	lib "./crypto"
)

// TODO: add argparse
// TODO: add sign
func main() {
	cert := lib.GetCertificate("examples/public.cer")
	keyStore := lib.GetKeyStore("examples/private.p12", "password")

	pubKey := lib.GetPublicKeyFromCert(cert)
	privKey := lib.GetPrivateKeyFromKeyStore(keyStore)

	secret := "Secret msg!"
	encrypted := lib.EncryptData(secret, pubKey)
	decrypted := lib.DecryptData(encrypted, privKey)

	fmt.Println("\nSecret msg: " + secret)
	fmt.Println("\nEncrypted msg: " + encrypted)
	fmt.Println("\nDecrypted msg: " + decrypted)
}
