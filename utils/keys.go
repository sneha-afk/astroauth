package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

var (
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
)

func LoadKeys(privPath string, pubPath string, keyType string) error {
	if keyType == "RSA" {
		if err := loadRSAPrivate(privPath); err != nil {
			return err
		}
		if err := loadRSAPublic(pubPath); err != nil {
			return err
		}
	}

	return nil
}

func loadRSAPrivate(path string) error {
	privFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("LoadKeys: failed opening private key file: %w", err)
	}
	defer privFile.Close()

	privKeyBytes, err := io.ReadAll(privFile)
	if err != nil {
		return fmt.Errorf("LoadKeys: failed reading private key file: %w", err)
	}

	privBlock, _ := pem.Decode(privKeyBytes)
	if privBlock == nil || privBlock.Type != "PRIVATE KEY" {
		return fmt.Errorf("LoadKeys: failed to decode private key PEM block")
	}

	loadedPrivKey, err := x509.ParsePKCS8PrivateKey(privBlock.Bytes)
	if err != nil {
		return fmt.Errorf("LoadKeys: could not parse private key: %w", err)
	}

	rsaPrivKey, ok := loadedPrivKey.(*rsa.PrivateKey)
	if !ok {
		return fmt.Errorf("LoadKeys: not an RSA private key")
	}

	PrivateKey = rsaPrivKey
	return nil
}

func loadRSAPublic(path string) error {
	pubFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("LoadKeys: failed opening public key file: %w", err)
	}
	defer pubFile.Close()

	pubKeyBytes, err := io.ReadAll(pubFile)
	if err != nil {
		return fmt.Errorf("LoadKeys: failed reading public key file: %w", err)
	}

	pubBlock, _ := pem.Decode(pubKeyBytes)
	if pubBlock == nil || pubBlock.Type != "PUBLIC KEY" {
		return fmt.Errorf("LoadKeys: failed to decode public key PEM block")
	}

	loadedPubKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return fmt.Errorf("LoadKeys: could not parse public key: %w", err)
	}

	rsaPubKey, ok := loadedPubKey.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("LoadKeys: not an RSA public key")
	}

	PublicKey = rsaPubKey
	return nil
}
