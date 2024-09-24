package util

import (
	"testing"
)

func TestNormalizeKey(t *testing.T) {
	tests := []struct {
		password    string
		expectedLen int
	}{
		{"short", 16},                           // 128-bit key
		{"thisislonger", 16},                    // 128-bit key
		{"thisisaverylongpassword", 24},         // 192-bit key
		{"thisisaveryveryverylongpassword", 32}, // 256-bit key
	}

	for _, test := range tests {
		key := NormalizeKey(test.password)
		if len(key) != test.expectedLen {
			t.Errorf("NormalizeKey(%q) = %d bytes; want %d bytes", test.password, len(key), test.expectedLen)
		}
	}
}
