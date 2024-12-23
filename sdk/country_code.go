package sdk

import (
	"fmt"
)

// CountryCode holds the country code tag and default country code.
type CountryCode struct {
	CountryCodeTag     string
	DefaultCountryCode string
}

// NewCountryCode initializes and returns a CountryCode instance.
func NewCountryCode(emv *EMV) *CountryCode {
	return &CountryCode{
		CountryCodeTag:     emv.CountryCode,        // Get CountryCode from the EMV struct
		DefaultCountryCode: emv.DefaultCountryCode, // Get DefaultCountryCode from the EMV struct
	}
}

// Value formats the country code according to the required structure.
func (c *CountryCode) Value(countryCode string) string {
	// Use the default if no country code is provided
	if countryCode == "" {
		countryCode = c.DefaultCountryCode
	}

	// Calculate the length of the country code
	lengthOfCountryCode := fmt.Sprintf("%02d", len(countryCode))

	// Construct and return the formatted result
	return fmt.Sprintf("%s%s%s", c.CountryCodeTag, lengthOfCountryCode, countryCode)
}
