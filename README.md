# ED25519 签名库

这是一个简单的Go语言ED25519数字签名实现，适用于区块链开发学习。

## 功能特点

- 生成ED25519密钥对
- 使用私钥签名消息
- 使用公钥验证签名
- 提供完整的单元测试
- 包含区块链交易签名的示例应用

## 项目结构

```
.
├── go.mod                  # Go模块定义文件
├── main.go                 # 示例应用
├── README.md               # 项目说明
└── signature/              # 签名库目录
    ├── ed25519.go          # 签名库主要实现
    └── ed25519_test.go     # 单元测试
```

## 使用方法

### 运行示例程序

```bash
go run main.go
```

### 运行单元测试

```bash
go test ./signature
```

## 在您的区块链项目中使用

```go
import "ed25519/signature"

// 生成密钥对
keyPair, err := signature.GenerateKeyPair()
if err != nil {
    // 处理错误
}

// 签名数据
message := []byte("需要签名的数据")
signature, err := keyPair.Sign(message)
if err != nil {
    // 处理错误
}

// 验证签名
isValid := keyPair.Verify(message, signature)
```

## 单元测试说明

单元测试是验证代码功能正确性的重要工具。本项目的测试涵盖了以下方面：

1. **密钥生成测试** - 确保能正确生成符合ED25519标准的密钥对
2. **签名与验证测试** - 确保签名后能被正确验证
3. **篡改消息测试** - 确保修改消息后签名失效
4. **错误密钥测试** - 确保使用错误的密钥无法验证签名

运行测试可以帮助您理解签名系统的安全特性，以及它在区块链中防止交易篡改的作用。 