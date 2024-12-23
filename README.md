# bakong-khqr (Unofficial)

A Go package for generating payment transactions compliant with the Bakong KHQR standard.

## Download Mobile App

- Bakong App ( [App Store](https://apps.apple.com/kh/app/bakong/id1440829141) | [Play Store](https://play.google.com/store/apps/details?id=jp.co.soramitsu.bakong) )
- Bakong Tourists ( [App Store](https://apps.apple.com/kh/app/bakong-tourists/id6471774666) | [Play Store](https://play.google.com/store/apps/details?id=kh.gov.nbc.bakong.tourist) )

## Requirement

- Go 1.23.3 or later
- Bakong Account (Full KYC)
- Bakong Developer Token [https://api-bakong.nbc.gov.kh/register](https://api-bakong.nbc.gov.kh/register)

## Installation

```bash
go get github.com/chhunneng/bakong-khqr
```

## Usage

The `bakong-khqr` package provides methods for generating QR codes, deeplinks, and verifying payment transactions using the Bakong KHQR standard.

### Importing the Package

```go
import "github.com/chhunneng/bakong-khqr"
```

### Creating a Payment Transaction

Example:

```go
package main

import (
    "fmt"
    "github.com/chhunneng/bakong-khqr"
)

func main() {
    // Initialize KHQR instance with Bakong Developer Token
    khqr := bakong_khqr.New("eyJhbGciOiJIUzI1NiIsI...nMhgG87BWeDg9Lu-_CKe1SMqC0")

    // Generate QR code data for a transaction
    qr, err := khqr.CreateQR(
        "your_name@wing", "Your Name", "Phnom Penh", 10000, "KHR", "MShop", "85512345678", "TRX019283775", "Cashier-01", false
    )
    if err != nil {
        fmt.Println("Error creating QR:", err)
        return
    }
    fmt.Println("QR Code:", qr)

    // Generate Deeplink
    deeplink, err := khqr.GenerateDeeplink(qr, "https://bakong.nbc.org.kh", "https://bakong.nbc.gov.kh/images/logo.svg", "YourApp")
    if err != nil {
        fmt.Println("Error generating deeplink:", err)
        return
    }
    fmt.Println("Deeplink:", deeplink)

    // Get MD5 Hash
    md5 := khqr.GenerateMD5(qr)
    fmt.Println("MD5 Hash:", md5)

    // Check Payment Status
    status, err := khqr.CheckPayment(md5)
    if err != nil {
        fmt.Println("Error checking payment:", err)
        return
    }
    fmt.Println("Payment Status:", status)
}
```

### Bulk Transaction Verification

To check multiple transactions:

```go
bulkStatus, err := khqr.CheckBulkPayments([]string{
    "dfcabf4598d1c405a75540a3d4ca099d",
    "5154e4f795634ff1a0ae4b48e53a6d9c",
    // Add more MD5 hashes as needed
})
if err != nil {
    fmt.Println("Error checking bulk payments:", err)
}
fmt.Println("Successful Transactions:", bulkStatus)
```

#### Parameters for `CreateQR` Method

- `bankAccount`: Associated bank account for the transaction.
- `merchantName`: Name of the merchant.
- `merchantCity`: City where the merchant is located.
- `amount`: Transaction amount.
- `currency`: Transaction currency (e.g., USD, KHR).
- `storeLabel`: Label or name of the store.
- `phoneNumber`: Merchant's contact number.
- `billNumber`: Reference number for the bill.
- `terminalLabel`: Terminal label for the transaction.
- `static`: Whether the QR is static or dynamic.

`Note`: Using static mode will create a Static QR Code for payment, allowing unlimited transactions, usage, and a zero amount included.

#### Parameters for `GenerateDeeplink` Method

- `qr`: QR code data generated by the CreateQR method.
- `callback`: Callback URL for post-payment redirection.
- `appIconUrl`: Icon URL for the app.
- `appName`: Application name.

## Bakong Official Documentation

- [https://api-bakong.nbc.gov.kh/document](https://api-bakong.nbc.gov.kh/document)
- [KHQR Content Guideline v1.4.pdf](https://bakong.nbc.gov.kh/download/KHQR/integration/KHQR%20Content%20Guideline%20v.1.3.pdf)
- [QR Payment Integration.pdf](https://bakong.nbc.gov.kh/download/KHQR/integration/QR%20Payment%20Integration.pdf)
- [KHQR SDK Document.pdf](https://bakong.nbc.gov.kh/download/KHQR/integration/KHQR%20SDK%20Document.pdf)

Development API: [https://sit-api-bakong.nbc.gov.kh/](https://sit-api-bakong.nbc.gov.kh/)

Production API: [https://api-bakong.nbc.gov.kh/](https://api-bakong.nbc.gov.kh/)

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/chhunneng/bakong-khqr/blob/main/LICENSE) file for details.

## Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request.
