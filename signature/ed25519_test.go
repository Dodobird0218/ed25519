package signature

import (
	"testing"
)

func TestGenerateKeyPair(t *testing.T) {
	keyPair, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Check if the key lengths conform to the ED25519 standard
	if len(keyPair.PrivateKey) != 64 {
		t.Errorf("Private key length error: expected 64 bytes, got %d bytes", len(keyPair.PrivateKey))
	}

	if len(keyPair.PublicKey) != 32 {
		t.Errorf("Public key length error: expected 32 bytes, got %d bytes", len(keyPair.PublicKey))
	}
}

func TestSignAndVerify(t *testing.T) {
	// Generate a key pair
	keyPair, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Test message
	message := []byte("Example message in blockchain application")

	// Sign the message
	signature, err := keyPair.Sign(message)
	if err != nil {
		t.Fatalf("Failed to sign message: %v", err)
	}

	// Verify the signature
	if !keyPair.Verify(message, signature) {
		t.Error("Signature verification failed, should pass")
	}

	// Use another verification method
	if !VerifyWithPublicKey(keyPair.PublicKey, message, signature) {
		t.Error("Signature verification using independent function failed, should pass")
	}
}

func TestVerifyWithTamperedMessage(t *testing.T) {
	// Generate a key pair
	keyPair, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Original message
	originalMessage := []byte("Original message in blockchain")

	// Tampered message
	tamperedMessage := []byte("Tampered message in blockchain")

	// Sign the original message
	signature, err := keyPair.Sign(originalMessage)
	if err != nil {
		t.Fatalf("Failed to sign message: %v", err)
	}

	// Verify with the tampered message, should fail
	if keyPair.Verify(tamperedMessage, signature) {
		t.Error("Verification succeeded with tampered message, should fail")
	}
}

func TestVerifyWithWrongKey(t *testing.T) {
	// Generate two key pairs
	keyPair1, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair 1: %v", err)
	}

	keyPair2, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair 2: %v", err)
	}

	// Test message
	message := []byte("Blockchain transaction information")

	// Sign with key pair 1
	signature, err := keyPair1.Sign(message)
	if err != nil {
		t.Fatalf("Failed to sign message: %v", err)
	}

	// Verify with key pair 2, should fail
	if keyPair2.Verify(message, signature) {
		t.Error("Verification succeeded with wrong key, should fail")
	}
}
