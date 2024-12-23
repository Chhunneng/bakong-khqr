package sdk

import (
	"fmt"
	"time"
)

// TimeStamp struct contains the logic for generating timestamp data
type TimeStamp struct {
	LanguagePreference string
	TimestampTag       string
}

// NewTimeStamp initializes and returns a new TimeStamp instance
func NewTimeStamp(emv *EMV) *TimeStamp {
	return &TimeStamp{
		LanguagePreference: emv.LanguagePreference,
		TimestampTag:       emv.TimestampTag,
	}
}

// Value generates the QR code data for the current timestamp
func (t *TimeStamp) Value() string {
	// Get the current timestamp in milliseconds
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())

	// Format the length of the timestamp
	lengthOfTimestamp := fmt.Sprintf("%02d", len(timestamp))

	// Create the initial result with language preference and timestamp
	result := fmt.Sprintf("%s%s%s", t.LanguagePreference, lengthOfTimestamp, timestamp)

	// Format the length of the result
	lengthResult := fmt.Sprintf("%02d", len(result))

	// Append the timestamp tag and formatted result
	return fmt.Sprintf("%s%s%s", t.TimestampTag, lengthResult, result)
}
