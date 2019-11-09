package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/pkcs12"
)

// public static byte[] encryptData(byte[] data,
// X509Certificate encryptionCertificate)
// throws CertificateEncodingException, CMSException, IOException {
// byte[] encryptedData = null;
// if (null != data && null != encryptionCertificate) {
// CMSEnvelopedDataGenerator cmsEnvelopedDataGenerator
// = new CMSEnvelopedDataGenerator();
// JceKeyTransRecipientInfoGenerator jceKey
// = new JceKeyTransRecipientInfoGenerator(encryptionCertificate);
// cmsEnvelopedDataGenerator.addRecipientInfoGenerator(jceKey);
// CMSTypedData msg = new CMSProcessableByteArray(data);
// OutputEncryptor encryptor
// = new JceCMSContentEncryptorBuilder(CMSAlgorithm.AES128_CBC)
// .setProvider("BC").build();
// CMSEnvelopedData cmsEnvelopedData = cmsEnvelopedDataGenerator
// .generate(msg,encryptor);
// encryptedData = cmsEnvelopedData.getEncoded();
// }
// return encryptedData;
// }

func encryptData(cer x509.Certificate, keyStore pkcs12.KeyStore) {

}

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

	// Create Private PKCS12 KeyStore
	p12, _ := ioutil.ReadFile("examples/private.p12")
	blocks, err := pkcs12.ToPEM(p12, string(keyPassword))
	if err != nil {
		panic(err)
	}
	fmt.Println(blocks)

}
