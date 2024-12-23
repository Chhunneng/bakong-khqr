package sdk

import (
	"fmt"
	"strconv"
	"strings"
)

// Amount holds the transaction amount tag and its maximum length.
type Amount struct {
	TransactionAmount string
	MaxLength         int
}

// NewAmount initializes and returns an Amount instance.
func NewAmount(emv *EMV) *Amount {
	return &Amount{
		TransactionAmount: emv.TransactionAmount,
		MaxLength:         emv.InvalidLengthAmount,
	}
}

// Value formats the amount according to the required structure.
func (a *Amount) Value(amount interface{}) (string, error) {
	// Ensure the amount is a valid type
	var amountStr string
	switch v := amount.(type) {
	case int:
		amountStr = fmt.Sprintf("%d", v)
	case float64:
		amountStr = fmt.Sprintf("%.2f", v)
	case string:
		_, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return "", fmt.Errorf("invalid amount value: %s. Amount must be a number or a string representing a number", v)
		}
		amountStr = v
	default:
		return "", fmt.Errorf("amount must be a number or a string")
	}

	// Remove trailing zeros after the decimal point
	if strings.Contains(amountStr, ".") {
		amountStr = strings.TrimRight(amountStr, "0")
		amountStr = strings.TrimRight(amountStr, ".")
	}

	// Ensure the length of the formatted amount does not exceed the max length
	lengthOfAmount := len(amountStr) + 2 // Adding 2 for the transaction amount tag and length
	if lengthOfAmount > a.MaxLength {
		return "", fmt.Errorf("formatted amount exceeds maximum length of %d characters. Your input length: %d characters", a.MaxLength, lengthOfAmount)
	}

	// Pad the amount to fit the required length
	paddedAmountStr := fmt.Sprintf("%011s", amountStr) // Pad with leading zeros

	// Calculate the length of the formatted amount string
	lengthOfAmountStr := fmt.Sprintf("%02d", len(paddedAmountStr))

	// Return the formatted amount string with the tag, length, and amount
	return fmt.Sprintf("%s%s%s", a.TransactionAmount, lengthOfAmountStr, paddedAmountStr), nil
}
