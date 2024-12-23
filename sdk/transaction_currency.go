package sdk

import (
	"fmt"
	"strings"
)

// TransactionCurrency struct contains the logic for generating transaction currency data
type TransactionCurrency struct {
	TransactionCurrency string
	CurrencyUSD         string
	CurrencyKHR         string
}

// NewTransactionCurrency initializes and returns a new TransactionCurrency instance
func NewTransactionCurrency(emv *EMV) *TransactionCurrency {
	return &TransactionCurrency{
		TransactionCurrency: emv.TransactionCurrency,
		CurrencyUSD:         emv.TransactionCurrencyUSD,
		CurrencyKHR:         emv.TransactionCurrencyKHR,
	}
}

// Value generates the QR code data for the transaction currency
func (tc *TransactionCurrency) Value(currency string) (string, error) {
	// Normalize the currency input
	currency = strings.ToUpper(currency)

	var currencyValue string

	// Determine the currency value based on input
	switch currency {
	case "USD":
		currencyValue = tc.CurrencyUSD
	case "KHR":
		currencyValue = tc.CurrencyKHR
	default:
		return "", fmt.Errorf("invalid currency code '%s', supported codes are 'USD' and 'KHR'", currency)
	}

	// Format the length of the currency value
	lengthOfCurrency := fmt.Sprintf("%02d", len(currencyValue))

	// Construct and return the formatted result
	return fmt.Sprintf("%s%s%s", tc.TransactionCurrency, lengthOfCurrency, currencyValue), nil
}
