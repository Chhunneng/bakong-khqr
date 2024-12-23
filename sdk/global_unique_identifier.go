package sdk

import (
	"fmt"
)

// GlobalUniqueIdentifier holds the values for payload format indicator, merchant account information, and max length.
type GlobalUniqueIdentifier struct {
	PayloadFormatIndicator               string
	MerchantAccountInformationIndividual string
	MaxLength                            int
}

// NewGlobalUniqueIdentifier initializes and returns a new GlobalUniqueIdentifier instance.
func NewGlobalUniqueIdentifier(emv *EMV) *GlobalUniqueIdentifier {
	return &GlobalUniqueIdentifier{
		PayloadFormatIndicator:               emv.PayloadFormatIndicator,
		MerchantAccountInformationIndividual: emv.MerchantAccountInformationIndividual,
		MaxLength:                            emv.InvalidLengthBakongAccount,
	}
}

// Value generates the global unique identifier based on the bank account number.
func (g *GlobalUniqueIdentifier) Value(bankAccount string) (string, error) {

	// Ensure the bank account does not exceed the maximum allowed length
	lengthOfBankAccount := len(bankAccount)
	if lengthOfBankAccount > g.MaxLength {
		return "", fmt.Errorf("bank account cannot exceed %d characters, your input length: %d characters", g.MaxLength, lengthOfBankAccount)
	}

	// Format the length of the bank account as two digits
	lengthOfBankAccountStr := fmt.Sprintf("%02d", lengthOfBankAccount)

	// Generate the result string
	result := fmt.Sprintf("%s%s%s", g.PayloadFormatIndicator, lengthOfBankAccountStr, bankAccount)

	// Format the length of the result as two digits
	lengthResult := fmt.Sprintf("%02d", len(result))

	// Final result
	return fmt.Sprintf("%s%s%s", g.MerchantAccountInformationIndividual, lengthResult, result), nil
}
