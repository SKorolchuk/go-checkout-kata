package history

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	underTest := NewScanHistory()

	assert.NoError(t, underTest.Add('A'))
}

func TestClean(t *testing.T) {
	underTest := NewScanHistory()

	assert.NoError(t, underTest.Clean())
}

func TestGetTotalUnitsPerSKU(t *testing.T) {
	expectedResult := map[rune]int32{
		'A': 2,
		'B': 1,
	}

	underTest := NewScanHistory()

	assert.NoError(t, underTest.Add('A'))
	assert.NoError(t, underTest.Add('A'))
	assert.NoError(t, underTest.Add('B'))

	actualResult := underTest.GetTotalUnitsPerSKU()

	assert.Equal(t, expectedResult, actualResult)
}
