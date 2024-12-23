package sdk

import (
	"fmt"
)

// AdditionalDataField holds the configuration for additional data fields based on the EMV configuration.
type AdditionalDataField struct {
	AdditionalDataTag   string
	StoreLabelTag       string
	MobileNumberTag     string
	BillNumberTag       string
	TerminalLabelTag    string
	StoreLabelLength    int
	MobileNumberLength  int
	BillNumberLength    int
	TerminalLabelLength int
}

// NewAdditionalDataField initializes and returns an AdditionalDataField instance with EMV configurations.
func NewAdditionalDataField(emv *EMV) *AdditionalDataField {
	return &AdditionalDataField{
		AdditionalDataTag:   emv.AdditionalDataTag,
		StoreLabelTag:       emv.StoreLabel,
		MobileNumberTag:     emv.AdditionDataFieldMobileNumber,
		BillNumberTag:       emv.BillNumberTag,
		TerminalLabelTag:    emv.TerminalLabel,
		StoreLabelLength:    emv.InvalidLengthStoreLabel,
		MobileNumberLength:  emv.InvalidLengthMobileNumber,
		BillNumberLength:    emv.InvalidLengthBillNumber,
		TerminalLabelLength: emv.InvalidLengthTerminalLabel,
	}
}

// formatValue formats a tag-value pair with a length prefix.
func (a *AdditionalDataField) formatValue(tag, value string) string {
	lengthOfValue := fmt.Sprintf("%02d", len(value))
	return tag + lengthOfValue + value
}

// validateLength checks if a value exceeds the maximum allowed length and raises an error if it does.
func (a *AdditionalDataField) validateLength(value string, maxLength int, fieldName string) error {
	if len(value) > maxLength {
		return fmt.Errorf("%s cannot exceed %d characters. Your input length: %d characters", fieldName, maxLength, len(value))
	}
	return nil
}

// StoreLabelValue formats and validates the store label value.
func (a *AdditionalDataField) StoreLabelValue(storeLabel string) (string, error) {
	if err := a.validateLength(storeLabel, a.StoreLabelLength, "Store label"); err != nil {
		return "", err
	}
	return a.formatValue(a.StoreLabelTag, storeLabel), nil
}

// PhoneNumberValue formats and validates the phone number value.
func (a *AdditionalDataField) PhoneNumberValue(phoneNumber string) (string, error) {
	if err := a.validateLength(phoneNumber, a.MobileNumberLength, "Phone number"); err != nil {
		return "", err
	}
	return a.formatValue(a.MobileNumberTag, phoneNumber), nil
}

// BillNumberValue formats and validates the bill number value.
func (a *AdditionalDataField) BillNumberValue(billNumber string) (string, error) {
	if err := a.validateLength(billNumber, a.BillNumberLength, "Bill number"); err != nil {
		return "", err
	}
	return a.formatValue(a.BillNumberTag, billNumber), nil
}

// TerminalLabelValue formats and validates the terminal label value.
func (a *AdditionalDataField) TerminalLabelValue(terminalLabel string) (string, error) {
	if err := a.validateLength(terminalLabel, a.TerminalLabelLength, "Terminal label"); err != nil {
		return "", err
	}
	return a.formatValue(a.TerminalLabelTag, terminalLabel), nil
}

// Value combines all formatted values into a single string with a length prefix.
func (a *AdditionalDataField) Value(storeLabel, phoneNumber, billNumber, terminalLabel string) (string, error) {
	storeLabelValue, err := a.StoreLabelValue(storeLabel)
	if err != nil {
		return "", err
	}
	phoneNumberValue, err := a.PhoneNumberValue(phoneNumber)
	if err != nil {
		return "", err
	}
	billNumberValue, err := a.BillNumberValue(billNumber)
	if err != nil {
		return "", err
	}
	terminalLabelValue, err := a.TerminalLabelValue(terminalLabel)
	if err != nil {
		return "", err
	}

	combinedData := billNumberValue + phoneNumberValue + storeLabelValue + terminalLabelValue
	lengthOfCombinedData := fmt.Sprintf("%02d", len(combinedData))

	return a.AdditionalDataTag + lengthOfCombinedData + combinedData, nil
}
