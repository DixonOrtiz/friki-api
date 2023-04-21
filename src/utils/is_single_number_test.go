package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	isSingle := IsSingleNumber(3)
	assert.True(t, isSingle)
}

func TestZeroSingleNumber(t *testing.T) {
	isSingle := IsSingleNumber(0)
	assert.True(t, isSingle)
}

func TestNonSingleNumber(t *testing.T) {
	isSingle := IsSingleNumber(10)
	assert.False(t, isSingle)
}
