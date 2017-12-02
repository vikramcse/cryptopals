package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
)

// the input is in hex format
// steps to convert hex -> base64
// 1. convert hex string format to decimal format
// 2. use the returned byte slice of decimal to convert onto base64
func hexToBase64(hexString string) (string, error) {
	s, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(s), nil
}
