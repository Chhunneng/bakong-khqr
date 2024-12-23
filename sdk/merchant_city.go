package sdk

import (
	"errors"
	"fmt"
)

// MerchantCity struct contains the merchant city logic
type MerchantCity struct {
	MerchantCityTag string
	MaxLength       int
}

// NewMerchantCity initializes and returns a new MerchantCity instance
func NewMerchantCity(emv *EMV) *MerchantCity {
	return &MerchantCity{
		MerchantCityTag: emv.MerchantCity,
		MaxLength:       emv.InvalidLengthMerchantCity,
	}
}

// Value generates and returns the formatted merchant city value
func (m *MerchantCity) Value(merchantCity string) (string, error) {
	// Validate the merchant city
	if merchantCity == "" {
		return "", errors.New("merchant city cannot be empty")
	}

	// Ensure the merchant city does not exceed the maximum allowed length
	if len(merchantCity) > m.MaxLength {
		return "", fmt.Errorf("merchant city cannot exceed %d characters. Your input length: %d characters", m.MaxLength, len(merchantCity))
	}

	// Calculate the length of the merchant city
	lengthStr := fmt.Sprintf("%02d", len(merchantCity))

	// Construct and return the result string
	return fmt.Sprintf("%s%s%s", m.MerchantCityTag, lengthStr, merchantCity), nil
}
