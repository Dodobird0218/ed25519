# ED25519 簽名庫

這是一個簡單的Go語言ED25519數字簽名實現，適用於區塊鏈開發學習。

## 功能特點

- 生成ED25519密鑰對
- 使用私鑰簽名消息
- 使用公鑰驗證簽名
- 提供完整的單元測試
- 包含區塊鏈交易簽名的示例應用

## 專案結構

```
.
├── go.mod                  # Go模組定義文件
├── main.go                 # 示例應用
├── README.md               # 專案說明
└── signature/              # 簽名庫目錄
    ├── ed25519.go          # 簽名庫主要實現
    └── ed25519_test.go     # 單元測試
```

## 使用方法

### 運行示例程序

```bash
go run main.go
```

### 運行單元測試

```bash
go test ./signature
```

## 在您的區塊鏈專案中使用

```go
import "ed25519/signature"

// 生成密鑰對
keyPair, err := signature.GenerateKeyPair()
if err != nil {
    // 處理錯誤
}

// 簽名數據
message := []byte("需要簽名的數據")
signature, err := keyPair.Sign(message)
if err != nil {
    // 處理錯誤
}

// 驗證簽名
isValid := keyPair.Verify(message, signature)
```

## 單元測試說明

單元測試是驗證代碼功能正確性的重要工具。本專案的測試涵蓋了以下方面：

1. **密鑰生成測試** - 確保能正確生成符合ED25519標準的密鑰對
2. **簽名與驗證測試** - 確保簽名後能被正確驗證
3. **篡改消息測試** - 確保修改消息後簽名失效
4. **錯誤密鑰測試** - 確保使用錯誤的密鑰無法驗證簽名

運行測試可以幫助您理解簽名系統的安全特性，以及它在區塊鏈中防止交易篡改的作用。 