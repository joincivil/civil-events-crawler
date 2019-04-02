package bytes

import (
	"encoding/hex"
)

// Byte32ToHexString converts a [32]byte slice to a string
func Byte32ToHexString(input [32]byte) string {
	return hex.EncodeToString(input[:])
}

// HexStringToByte32 converts a string back to a [32]byte slice
func HexStringToByte32(input string) ([32]byte, error) {
	bys, err := hex.DecodeString(input)
	fixed := [32]byte{}
	if err != nil {
		return fixed, err
	}
	copy(fixed[:], bys)
	return fixed, nil
}
