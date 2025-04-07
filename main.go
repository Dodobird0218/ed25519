package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"ed25519/signature"
)

// 模拟区块链交易结构
type Transaction struct {
	Sender    string
	Receiver  string
	Amount    float64
	Timestamp int64
}

// 将交易结构转换为字节数组用于签名
func (tx *Transaction) ToBytes() []byte {
	// 实际应用中应使用更规范的序列化方法，如Protobuf或JSON
	data := fmt.Sprintf("%s-%s-%.2f-%d", tx.Sender, tx.Receiver, tx.Amount, tx.Timestamp)
	return []byte(data)
}

func main() {
	// 生成一对新的密钥
	keyPair, err := signature.GenerateKeyPair()
	if err != nil {
		log.Fatalf("生成密钥对失败: %v", err)
	}

	// 打印密钥信息
	fmt.Printf("生成的公钥: %s\n", hex.EncodeToString(keyPair.PublicKey))
	fmt.Printf("生成的私钥: %s\n\n", hex.EncodeToString(keyPair.PrivateKey))

	// 创建一个模拟的区块链交易
	transaction := &Transaction{
		Sender:    "Alice",
		Receiver:  "Bob",
		Amount:    1.5,
		Timestamp: 1680153600,
	}

	// 获取交易数据并签名
	txData := transaction.ToBytes()
	signature, err := keyPair.Sign(txData)
	if err != nil {
		log.Fatalf("签名交易失败: %v", err)
	}

	fmt.Printf("交易数据: %s\n", string(txData))
	fmt.Printf("签名: %s\n\n", hex.EncodeToString(signature))

	// 验证签名
	isValid := keyPair.Verify(txData, signature)
	fmt.Printf("签名验证结果: %v\n", isValid)

	// 模拟篡改交易数据
	tamperedTransaction := &Transaction{
		Sender:    "Alice",
		Receiver:  "Bob",
		Amount:    10.5, // 改变金额
		Timestamp: 1680153600,
	}

	tamperedTxData := tamperedTransaction.ToBytes()
	fmt.Printf("\n篡改后的交易数据: %s\n", string(tamperedTxData))

	// 验证篡改后的交易
	isValidTampered := keyPair.Verify(tamperedTxData, signature)
	fmt.Printf("篡改后的交易验证结果: %v\n", isValidTampered)
}
