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

## ED25519 原理與公式

ED25519 是一種基於椭圆曲线的數字簽名算法，使用的是 Curve25519 椭圆曲线。它具有高效性和安全性，廣泛應用於區塊鏈和其他安全通信領域。

### 基本原理

ED25519 的簽名過程主要包括以下幾個步驟：

1. **密鑰生成**：
   - 生成一對密鑰：私鑰和公鑰。私鑰是隨機生成的，而公鑰是通過椭圆曲线的運算從私鑰導出的。

2. **簽名**：
   - 將要簽名的消息進行哈希處理，生成固定長度的摘要。
   - 使用私鑰和消息摘要生成簽名，這個過程涉及到椭圆曲线的運算。

3. **驗證**：
   - 接收方使用公鑰和簽名來驗證消息的完整性和來源。這個過程同樣涉及椭圆曲线的運算。

### 數學公式

ED25519 使用的主要數學公式包括：

- **椭圆曲线方程**：
  \[
  y^2 = x^3 + 486662x^2 + x
  \]
  這是 Curve25519 的基本方程。

- **簽名生成**：
  簽名 \( (R, S) \) 的生成過程可以表示為：
  \[
  R = k \cdot B
  \]
  \[
  S = (H(m || R) + a \cdot R) \mod L
  \]
  其中 \( k \) 是隨機數，\( B \) 是基點，\( H \) 是哈希函數，\( a \) 是私鑰，\( L \) 是曲線的階。

- **簽名驗證**：
  驗證過程可以表示為：
  \[
  H(m || R) \cdot B = S \cdot B + R \cdot A
  \]
  其中 \( A \) 是公鑰。

這些公式和過程確保了 ED25519 簽名的安全性和有效性，並防止了常見的攻擊方式，如重放攻擊和篡改攻擊。

## 單元測試說明

單元測試是驗證代碼功能正確性的重要工具。本專案的測試涵蓋了以下方面：

1. **密鑰生成測試** - 確保能正確生成符合ED25519標準的密鑰對
2. **簽名與驗證測試** - 確保簽名後能被正確驗證
3. **篡改消息測試** - 確保修改消息後簽名失效
4. **錯誤密鑰測試** - 確保使用錯誤的密鑰無法驗證簽名

運行測試可以幫助您理解簽名系統的安全特性，以及它在區塊鏈中防止交易篡改的作用。 