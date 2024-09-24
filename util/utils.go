package util

import (
	"crypto/sha256"
)

// NormalizeKey ensures the key is 16 bytes (AES-128), 24 bytes (AES-192), or 32 bytes (AES-256)
func NormalizeKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	key := hash[:]

	switch {
	case len(password) <= 16:
		return key[:16] // 128-bit key
	case len(password) <= 24:
		return key[:24] // 192-bit key
	default:
		return key[:32] // 256-bit key
	}
}
