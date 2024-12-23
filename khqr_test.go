package khqr

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func setupKHQR() *KHQR {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bakongToken := os.Getenv("BAKONG_TOKEN")
	// Initialize KHQR with Bakong Token
	khqrInstance := NewKHQR(bakongToken)
	return khqrInstance
}

func TestCreateQR(t *testing.T) {
	khqrInstance := setupKHQR()

	// Create a QR code string
	qr, err := khqrInstance.CreateQR("your_name@wing", "Your Name", "Phnom Penh", 10000, "KHR", "MShop", "85512345678", "TRX019283775", "Cashier-01", false)
	if err != nil {
		t.Fatalf("Failed to create QR: %v", err)
	}

	// Generate Deeplink
	deeplink, err := khqrInstance.GenerateDeeplink(qr, "https://bakong.nbc.org.kh", "https://bakong.nbc.gov.kh/images/logo.svg", "MyAppName")
	if err != nil {
		t.Fatalf("Failed to generate deeplink: %v", err)
	}

	// Generate MD5 hash
	md5 := khqrInstance.GenerateMD5(qr)

	// Check transaction status
	paymentStatus, err := khqrInstance.CheckPayment(md5)
	if err != nil {
		t.Fatalf("Failed to check payment status: %v", err)
	}

	// Check bulk transactions
	md5List := []string{
		"dfcabf4598d1c405a75540a3d4ca099d",
		"5154e4f795634ff1a0ae4b48e53a6d9c",
		"a57d9bb85f52f12a20cf7beecb03d11d",
		"495fdaec0be5d94c89bc1283c7283d3d",
		"31bca02094ad576588e42b60db57bc98",
	}
	bulkPaymentsStatus, err := khqrInstance.CheckBulkPayments(md5List)
	if err != nil {
		t.Fatalf("Failed to check bulk payment statuses: %v", err)
	}

	// Print the results
	t.Logf("QR Code Data: %s", qr)
	t.Logf("Deeplink: %s", deeplink)
	t.Logf("Transaction MD5: %s", md5)
	t.Logf("Check Payment Status: %v", paymentStatus)
	t.Logf("Check Bulk Payments Status: %v", bulkPaymentsStatus)
}
