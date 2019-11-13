package crypto

import (
	"errors"
	"testing"
)

func TestEcryption(t *testing.T) {

	cert := GetCertificate("public.cer")
	keyStore := GetKeyStore("private.p12", "password")

	pubKey := GetPublicKeyFromCert(cert)
	privKey := GetPrivateKeyFromKeyStore(keyStore)

	secret := "Secret msg!"
	encrypted := EncryptData(secret, pubKey)
	decrypted := DecryptData(encrypted, privKey)

	checkError(errors.New(""))

	if secret != decrypted {
		t.Errorf("Decryption was incorrect, got: %s, want: %s.", decrypted, secret)
	}
}
