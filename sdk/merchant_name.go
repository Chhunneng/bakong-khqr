package sdk

import (
	"errors"
	"fmt"
)

// MerchantName struct contains the merchant name logic
type MerchantName struct {
	MerchantNameTag string
	MaxLength       int
}

// NewMerchantName initializes and returns a new MerchantName instance
func NewMerchantName(emv *EMV) *MerchantName {
	return &MerchantName{
		MerchantNameTag: emv.MerchantName,
		MaxLength:       emv.InvalidLengthMerchantName,
	}
}

// Value generates and returns the formatted merchant name value
func (m *MerchantName) Value(merchantName string) (string, error) {
	// Validate the merchant name
	if merchantName == "" {
		return "", errors.New("merchant name cannot be empty")
	}

	// Ensure the merchant name does not exceed the maximum allowed length
	if len(merchantName) > m.MaxLength {
		return "", fmt.Errorf("merchant name cannot exceed %d characters. Your input length: %d characters", m.MaxLength, len(merchantName))
	}

	// Calculate the length of the merchant name
	lengthStr := fmt.Sprintf("%02d", len(merchantName))

	// Construct and return the result string
	return fmt.Sprintf("%s%s%s", m.MerchantNameTag, lengthStr, merchantName), nil
}
