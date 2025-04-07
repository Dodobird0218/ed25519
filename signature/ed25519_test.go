package signature

import (
	"testing"
)

func TestGenerateKeyPair(t *testing.T) {
	keyPair, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 检查密钥长度是否符合ED25519标准
	if len(keyPair.PrivateKey) != 64 {
		t.Errorf("私钥长度错误: 期望64字节, 实际%d字节", len(keyPair.PrivateKey))
	}

	if len(keyPair.PublicKey) != 32 {
		t.Errorf("公钥长度错误: 期望32字节, 实际%d字节", len(keyPair.PublicKey))
	}
}

func TestSignAndVerify(t *testing.T) {
	// 生成密钥对
	keyPair, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 测试消息
	message := []byte("区块链应用中的消息示例")

	// 签名消息
	signature, err := keyPair.Sign(message)
	if err != nil {
		t.Fatalf("签名消息失败: %v", err)
	}

	// 验证签名
	if !keyPair.Verify(message, signature) {
		t.Error("签名验证失败, 应该验证通过")
	}

	// 使用另一种验证方法
	if !VerifyWithPublicKey(keyPair.PublicKey, message, signature) {
		t.Error("使用独立函数验证签名失败, 应该验证通过")
	}
}

func TestVerifyWithTamperedMessage(t *testing.T) {
	// 生成密钥对
	keyPair, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 原始消息
	originalMessage := []byte("区块链中的原始消息")

	// 篡改后的消息
	tamperedMessage := []byte("区块链中的篡改消息")

	// 对原始消息签名
	signature, err := keyPair.Sign(originalMessage)
	if err != nil {
		t.Fatalf("签名消息失败: %v", err)
	}

	// 用篡改后的消息验证, 应该失败
	if keyPair.Verify(tamperedMessage, signature) {
		t.Error("篡改消息后验证成功, 应该验证失败")
	}
}

func TestVerifyWithWrongKey(t *testing.T) {
	// 生成两对密钥对
	keyPair1, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("生成密钥对1失败: %v", err)
	}

	keyPair2, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("生成密钥对2失败: %v", err)
	}

	// 测试消息
	message := []byte("区块链交易信息")

	// 用密钥对1签名
	signature, err := keyPair1.Sign(message)
	if err != nil {
		t.Fatalf("签名消息失败: %v", err)
	}

	// 用密钥对2验证, 应该失败
	if keyPair2.Verify(message, signature) {
		t.Error("使用错误的密钥验证成功, 应该验证失败")
	}
}
