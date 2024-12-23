package sdk

import (
	"errors"
	"fmt"
)

// MCC struct contains the Merchant Category Code (MCC) logic
type MCC struct {
	MerchantCategoryCodeTag     string
	DefaultMerchantCategoryCode string
}

// NewMCC initializes and returns a new MCC instance
func NewMCC(emv *EMV) *MCC {
	return &MCC{
		MerchantCategoryCodeTag:     emv.MerchantCategoryCode,
		DefaultMerchantCategoryCode: emv.DefaultMerchantCategoryCode,
	}
}

// Value generates and returns the formatted Merchant Category Code (MCC) value
func (m *MCC) Value(categoryCode string) (string, error) {
	// Use the default category code if none is provided
	if categoryCode == "" {
		categoryCode = m.DefaultMerchantCategoryCode
	}

	// Validate the category code: must be numeric and at least 4 digits long
	if len(categoryCode) < 4 || !isNumeric(categoryCode) {
		return "", errors.New("category code must be a numeric string with at least 4 digits")
	}

	// Calculate the length of the category code
	lengthStr := fmt.Sprintf("%02d", len(categoryCode))

	// Construct and return the result string
	return fmt.Sprintf("%s%s%s", m.MerchantCategoryCodeTag, lengthStr, categoryCode), nil
}

// isNumeric checks if a string is numeric
func isNumeric(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}
