package utils

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

var (
	PublicKey  any
	PrivateKey any
)

func LoadKeys(privPath string, pubPath string, keyType string) error {
	// Get bytes from both private and public key files
	privFile, err := os.Open(privPath)
	if err != nil {
		return fmt.Errorf("LoadKeys: failed opening private key file: %w", err)
	}
	defer privFile.Close()

	privKeyBytes, err := io.ReadAll(privFile)
	if err != nil {
		return fmt.Errorf("LoadKeys: failed reading private key file: %w", err)
	}

	pubFile, err := os.Open(pubPath)
	if err != nil {
		return fmt.Errorf("LoadKeys: failed opening public key file: %w", err)
	}
	defer pubFile.Close()

	pubKeyBytes, err := io.ReadAll(pubFile)
	if err != nil {
		return fmt.Errorf("LoadKeys: failed reading public key file: %w", err)
	}

	// Parse according to signing method
	if keyType == "RSA" {
		if err := loadRSAPrivate(privKeyBytes); err != nil {
			return err
		}
		if err := loadRSAPublic(pubKeyBytes); err != nil {
			return err
		}
	} else if keyType == "ECDSA" {
		if err := loadECDSAPrivate(privKeyBytes); err != nil {
			return err
		}
		if err := loadECDSAPublic(pubKeyBytes); err != nil {
			return err
		}
	}

	return nil
}

func loadRSAPrivate(privBytes []byte) error {
	privBlock, _ := pem.Decode(privBytes)
	if privBlock == nil || privBlock.Type != "PRIVATE KEY" {
		return fmt.Errorf("loadRSAPrivate: failed to decode private key PEM block")
	}

	loadedPrivKey, err := x509.ParsePKCS8PrivateKey(privBlock.Bytes)
	if err != nil {
		return fmt.Errorf("loadRSAPrivate: could not parse private key: %w", err)
	}

	rsaPrivKey, ok := loadedPrivKey.(*rsa.PrivateKey)
	if !ok {
		return fmt.Errorf("loadRSAPrivate: not an RSA private key")
	}

	PrivateKey = rsaPrivKey
	return nil
}

func loadRSAPublic(pubBytes []byte) error {
	pubBlock, _ := pem.Decode(pubBytes)
	if pubBlock == nil || pubBlock.Type != "PUBLIC KEY" {
		return fmt.Errorf("loadRSAPublic: failed to decode public key PEM block")
	}

	loadedPubKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return fmt.Errorf("loadRSAPublic: could not parse public key: %w", err)
	}

	rsaPubKey, ok := loadedPubKey.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("loadRSAPublic: not an RSA public key")
	}

	PublicKey = rsaPubKey
	return nil
}

func loadECDSAPrivate(privBytes []byte) error {
	loadedPrivKey, err := x509.ParseECPrivateKey(privBytes)
	if err != nil {
		return fmt.Errorf("loadECDSAPrivate: could not parse private key: %w", err)
	}

	PrivateKey = loadedPrivKey
	return nil
}

func loadECDSAPublic(pubBytes []byte) error {
	pubBlock, _ := pem.Decode(pubBytes)
	if pubBlock == nil || pubBlock.Type != "PUBLIC KEY" {
		return fmt.Errorf("loadECDSAPublic: failed to decode public key PEM block")
	}

	loadedPubKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return fmt.Errorf("loadECDSAPublic: could not parse public key: %w", err)
	}

	ecdsaPubKey, ok := loadedPubKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("loadECDSAPublic: not an RSA public key")
	}

	PublicKey = ecdsaPubKey
	return nil
}
