package signature

import (
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"fmt"
)

// KeyPair 包含ED25519的公钥和私钥
type KeyPair struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
}

// GenerateKeyPair 生成一个新的ED25519密钥对
func GenerateKeyPair() (*KeyPair, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("生成密钥对失败: %w", err)
	}

	return &KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

// Sign 使用私钥对消息进行签名
func (kp *KeyPair) Sign(message []byte) ([]byte, error) {
	if kp.PrivateKey == nil {
		return nil, errors.New("私钥不存在")
	}

	signature := ed25519.Sign(kp.PrivateKey, message)
	return signature, nil
}

// Verify 使用公钥验证消息签名
func (kp *KeyPair) Verify(message, signature []byte) bool {
	if kp.PublicKey == nil {
		return false
	}

	return ed25519.Verify(kp.PublicKey, message, signature)
}

// VerifyWithPublicKey 使用提供的公钥验证消息签名
func VerifyWithPublicKey(publicKey ed25519.PublicKey, message, signature []byte) bool {
	return ed25519.Verify(publicKey, message, signature)
}
