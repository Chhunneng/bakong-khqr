package khqr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/chhunneng/bakong-khqr/sdk"
)

// Define the KHQR struct
type KHQR struct {
	crc                    sdk.CRC
	mcc                    sdk.MCC
	hash                   sdk.HASH
	amount                 sdk.Amount
	timestamp              sdk.TimeStamp
	countryCode            sdk.CountryCode
	merchantCity           sdk.MerchantCity
	merchantName           sdk.MerchantName
	pointOfInitiation      sdk.PointOfInitiation
	transactionCurrency    sdk.TransactionCurrency
	additionalDataField    sdk.AdditionalDataField
	payloadFormatIndicator sdk.PayloadFormatIndicator
	globalUniqueIdentifier sdk.GlobalUniqueIdentifier
	bakongToken            string
	bakongAPI              string
}

// Initialize the KHQR struct
func NewKHQR(bakongToken string) *KHQR {

	emv := sdk.NewEMV()

	return &KHQR{
		crc:                    *sdk.NewCRC(emv),
		mcc:                    *sdk.NewMCC(emv),
		hash:                   *sdk.NewHASH(),
		amount:                 *sdk.NewAmount(emv),
		timestamp:              *sdk.NewTimeStamp(emv),
		countryCode:            *sdk.NewCountryCode(emv),
		merchantCity:           *sdk.NewMerchantCity(emv),
		merchantName:           *sdk.NewMerchantName(emv),
		pointOfInitiation:      *sdk.NewPointOfInitiation(emv),
		transactionCurrency:    *sdk.NewTransactionCurrency(emv),
		additionalDataField:    *sdk.NewAdditionalDataField(emv),
		payloadFormatIndicator: *sdk.NewPayloadFormatIndicator(emv),
		globalUniqueIdentifier: *sdk.NewGlobalUniqueIdentifier(emv),
		bakongToken:            bakongToken,
		bakongAPI:              "https://api-bakong.nbc.gov.kh/v1",
	}
}

// Method to create QR code
func (khqr *KHQR) CreateQR(bankAccount string, merchantName string, merchantCity string, amount float64, currency string, storeLabel string, phoneNumber string, billNumber string, terminalLabel string, static bool) (string, error) {
	qrData := khqr.payloadFormatIndicator.Value()
	if static {
		qrData += khqr.pointOfInitiation.Static()
	} else {
		qrData += khqr.pointOfInitiation.Dynamic()
	}
	result, err := khqr.globalUniqueIdentifier.Value(bankAccount)
	if err != nil {
		return "", err
	}
	qrData += result
	result, err = khqr.mcc.Value("")
	if err != nil {
		return "", err
	}
	qrData += result
	qrData += khqr.countryCode.Value("")
	result, err = khqr.merchantName.Value(merchantName)
	if err != nil {
		return "", err
	}
	qrData += result
	result, err = khqr.merchantCity.Value(merchantCity)
	if err != nil {
		return "", err
	}
	qrData += result
	qrData += khqr.timestamp.Value()
	if !static {
		result, err = khqr.amount.Value(amount)
		if err != nil {
			return "", err
		}
		qrData += result
	}
	result, err = khqr.transactionCurrency.Value(currency)
	if err != nil {
		return "", err
	}
	qrData += result
	result, err = khqr.additionalDataField.Value(storeLabel, phoneNumber, billNumber, terminalLabel)
	if err != nil {
		return "", err
	}
	qrData += result
	qrData += khqr.crc.Value(qrData)

	return qrData, nil
}

// Method to generate deep link
func (khqr *KHQR) GenerateDeeplink(qr string, callback string, appIconUrl string, appName string) (string, error) {

	if khqr.bakongToken == "" {
		return "", fmt.Errorf("bakong developer token is required for KHQR class initialization")
	}
	if callback == "" {
		callback = "https://bakong.nbc.org.kh"
	}
	if appIconUrl == "" {
		appIconUrl = "https://bakong.nbc.gov.kh/images/logo.svg"
	}
	if appName == "" {
		appName = "MyAppName"
	}

	payload := map[string]interface{}{
		"qr": qr,
		"sourceInfo": map[string]interface{}{
			"appIconUrl":          appIconUrl,
			"appName":             appName,
			"appDeepLinkCallback": callback,
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", khqr.bakongAPI+"/generate_deeplink_by_qr", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+khqr.bakongToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}
	if response["responseCode"].(float64) == 0 {
		return response["data"].(map[string]interface{})["shortLink"].(string), nil
	} else {
		return "", fmt.Errorf("error: %v", response["status"].(map[string]interface{})["message"])
	}
}

// Method to generate MD5 hash
func (khqr *KHQR) GenerateMD5(qr string) string {
	return khqr.hash.Md5(qr)
}

// Method to check payment status
func (khqr *KHQR) CheckPayment(md5 string) (string, error) {
	if khqr.bakongToken == "" {
		return "", fmt.Errorf("the Bakong Developer Token is required for KHQR class initialization")
	}

	payload := map[string]string{
		"md5": md5,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", khqr.bakongAPI+"/check_transaction_by_md5", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+khqr.bakongToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	if response["responseCode"] == 0 {
		return "PAID", nil
	} else if response["responseCode"] == 1 && response["errorCode"] == 6 {
		return "", fmt.Errorf("our developer token is either incorrect or expired, please renew it through Bakong Developer")
	}

	return "UNPAID", nil
}

// Method to check bulk payments
func (khqr *KHQR) CheckBulkPayments(md5List []string) ([]string, error) {
	if khqr.bakongToken == "" {
		return nil, fmt.Errorf("bakong developer token is required for KHQR class initialization")
	}

	payloadBytes, err := json.Marshal(md5List)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", khqr.bakongAPI+"/check_transaction_by_md5_list", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+khqr.bakongToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if response["responseCode"] == 0 {
		var paidList []string
		for _, data := range response["data"].([]interface{}) {
			if data.(map[string]interface{})["status"] == "SUCCESS" {
				paidList = append(paidList, data.(map[string]interface{})["md5"].(string))
			}
		}
		return paidList, nil
	} else if response["responseCode"] == 1 && response["errorCode"] == 6 {
		return nil, fmt.Errorf("your developer token is either incorrect or expired. please renew it through Bakong Developer")
	}

	return []string{}, nil
}
