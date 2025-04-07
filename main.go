package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"ed25519/signature"
)

// Simulated blockchain transaction structure
type Transaction struct {
	Sender    string
	Receiver  string
	Amount    float64
	Timestamp int64
}

// Convert transaction structure to byte array for signing
func (tx *Transaction) ToBytes() []byte {
	// In a real application, a more standardized serialization method should be used, such as Protobuf or JSON
	data := fmt.Sprintf("%s-%s-%.2f-%d", tx.Sender, tx.Receiver, tx.Amount, tx.Timestamp)
	return []byte(data)
}

func main() {
	// Generate a new key pair
	keyPair, err := signature.GenerateKeyPair()
	if err != nil {
		log.Fatalf("Failed to generate key pair: %v", err)
	}

	// Print key information
	fmt.Printf("Generated public key: %s\n", hex.EncodeToString(keyPair.PublicKey))
	fmt.Printf("Generated private key: %s\n\n", hex.EncodeToString(keyPair.PrivateKey))

	// Create a simulated blockchain transaction
	transaction := &Transaction{
		Sender:    "Alice",
		Receiver:  "Bob",
		Amount:    1.5,
		Timestamp: 1680153600,
	}

	// Get transaction data and sign it
	txData := transaction.ToBytes()
	signature, err := keyPair.Sign(txData)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	fmt.Printf("Transaction data: %s\n", string(txData))
	fmt.Printf("Signature: %s\n\n", hex.EncodeToString(signature))

	// Verify the signature
	isValid := keyPair.Verify(txData, signature)
	fmt.Printf("Signature verification result: %v\n", isValid)

	// Simulate tampering with the transaction data
	tamperedTransaction := &Transaction{
		Sender:    "Alice",
		Receiver:  "Bob",
		Amount:    10.5, // Change amount
		Timestamp: 1680153600,
	}

	tamperedTxData := tamperedTransaction.ToBytes()
	fmt.Printf("\nTampered transaction data: %s\n", string(tamperedTxData))

	// Verify the tampered transaction
	isValidTampered := keyPair.Verify(tamperedTxData, signature)
	fmt.Printf("Tampered transaction verification result: %v\n", isValidTampered)
}
