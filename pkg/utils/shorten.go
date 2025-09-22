package utils

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

// GenerateShortCode creates a random short string
func GenerateShortCode(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	// Encode to base64 and strip unwanted chars
	code := base64.URLEncoding.EncodeToString(b)
	code = strings.TrimRight(code, "=")
	if len(code) > length {
		code = code[:length]
	}
	return code, nil
}
