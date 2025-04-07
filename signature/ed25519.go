package signature

import (
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"fmt"
)

// KeyPair contains the public and private keys for ED25519
type KeyPair struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
}

// GenerateKeyPair generates a new ED25519 key pair
func GenerateKeyPair() (*KeyPair, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key pair: %w", err)
	}

	return &KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

// Sign signs a message using the private key
func (kp *KeyPair) Sign(message []byte) ([]byte, error) {
	if kp.PrivateKey == nil {
		return nil, errors.New("private key does not exist")
	}

	signature := ed25519.Sign(kp.PrivateKey, message)
	return signature, nil
}

// Verify verifies a message signature using the public key
func (kp *KeyPair) Verify(message, signature []byte) bool {
	if kp.PublicKey == nil {
		return false
	}

	return ed25519.Verify(kp.PublicKey, message, signature)
}

// VerifyWithPublicKey verifies a message signature using the provided public key
func VerifyWithPublicKey(publicKey ed25519.PublicKey, message, signature []byte) bool {
	return ed25519.Verify(publicKey, message, signature)
}
