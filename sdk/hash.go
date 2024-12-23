package sdk

import (
	"crypto/md5"
	"encoding/hex"
)

// HASH struct is an empty struct to represent the hashing functionality.
type HASH struct{}

// NewHASH initializes and returns a new HASH instance.
func NewHASH() *HASH {
	return &HASH{}
}

// Md5 generates an MD5 hash for the given data.
func (h *HASH) Md5(data string) string {

	// Create MD5 hash of the data
	hash := md5.New()
	hash.Write([]byte(data))

	// Return the hexadecimal representation of the MD5 hash
	return hex.EncodeToString(hash.Sum(nil))
}
