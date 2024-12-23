package sdk

import (
	"fmt"
)

// CRC holds the CRC tag and default CRC tag values.
type CRC struct {
	CRC           string
	DefaultCRCTag string
}

// NewCRC initializes and returns a CRC instance.
func NewCRC(emv *EMV) *CRC {
	return &CRC{
		CRC:           emv.CRC,           // Get CRC from the EMV struct
		DefaultCRCTag: emv.DefaultCRCTag, // Get Default CRC tag from the EMV struct
	}
}

// CalculateCRC16 calculates the CRC-16 using the CRC-CCITT polynomial.
func (c *CRC) CalculateCRC16(data string) int {
	crc := 0xFFFF        // Initial CRC value
	polynomial := 0x1021 // CRC-CCITT polynomial

	// Loop through each byte of data and calculate CRC
	for i := 0; i < len(data); i++ {
		byteData := data[i]
		crc ^= int(byteData) << 8

		for j := 0; j < 8; j++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ polynomial
			} else {
				crc <<= 1
			}
			crc &= 0xFFFF // Ensure it stays 16-bit
		}
	}
	return crc
}

// CRC16Hex returns the CRC-16 value in hexadecimal format.
func (c *CRC) CRC16Hex(data string) string {
	crc16Result := c.CalculateCRC16(data)
	return fmt.Sprintf("%04X", crc16Result)
}

// Value computes the CRC-16 value and formats it including the CRC tag.
func (c *CRC) Value(data string) string {
	// Calculate CRC-16 value including the default CRC tag
	crc16Hex := c.CRC16Hex(data + c.DefaultCRCTag)
	lengthOfCRC := fmt.Sprintf("%02d", len(crc16Hex)) // Calculate length of CRC in 2 digits

	// Return formatted string with CRC tag, length, and CRC value
	return fmt.Sprintf("%s%s%s", c.CRC, lengthOfCRC, crc16Hex)
}
