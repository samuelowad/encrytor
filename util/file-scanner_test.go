package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScanFiles(t *testing.T) {
	fileData, err := ScanFiles("./")
	require.NoError(t, err)
	require.NotEmpty(t, fileData)
}
