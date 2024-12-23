package sdk

import (
	"fmt"
)

// PayloadFormatIndicator struct contains the payload format indicator logic
type PayloadFormatIndicator struct {
	PayloadFormatIndicator        string
	DefaultPayloadFormatIndicator string
}

// NewPayloadFormatIndicator initializes and returns a new PayloadFormatIndicator instance
func NewPayloadFormatIndicator(emv *EMV) *PayloadFormatIndicator {
	return &PayloadFormatIndicator{
		PayloadFormatIndicator:        emv.PayloadFormatIndicator,
		DefaultPayloadFormatIndicator: emv.DefaultPayloadFormatIndicator,
	}
}

// Value generates and returns the formatted payload format indicator value
func (p *PayloadFormatIndicator) Value() string {
	// Calculate the length of the default payload format indicator
	length := len(p.DefaultPayloadFormatIndicator)

	// Format the length as two digits
	lengthStr := fmt.Sprintf("%02d", length)

	// Construct and return the result string
	return fmt.Sprintf("%s%s%s", p.PayloadFormatIndicator, lengthStr, p.DefaultPayloadFormatIndicator)
}
