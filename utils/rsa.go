package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func CreateKeyPair() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// Encode the public key to PEM format
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		fmt.Println("Error encoding public key:", err)
		return
	}
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	privateKeyPEM := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	fmt.Println("private: \n", string(pem.EncodeToMemory(privateKeyPEM)))
	fmt.Println("public: \n", string(pem.EncodeToMemory(publicKeyPEM)))
}

func DecryptWithPrivateKey(cipherText string) ([]byte, error) {
	privateKey, err := getPrivateKey()
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, []byte(cipherText))
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	// Read the private key from the PEM-encoded file
	privateKeyPEM, err := os.ReadFile("private.pem")
	if err != nil {
		return nil, fmt.Errorf("error reading private key file: %s", err)
	}

	block, _ := pem.Decode(privateKeyPEM)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	return key, nil
}

func GetPublicKey() (string, error) {
	// Read the public key from the PEM-encoded file
	publicKeyPEM, err := os.ReadFile("public.pem")
	if err != nil {
		return "", fmt.Errorf("error reading public key file: %s", err)
	}

	return string(publicKeyPEM), nil
}
