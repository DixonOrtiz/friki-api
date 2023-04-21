package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberExists(t *testing.T) {
	exists := DoesNumberExist(3, []int{1, 2, 3})
	assert.True(t, exists)
}

func TestNumberDoesNotExist(t *testing.T) {
	exists := DoesNumberExist(4, []int{1, 2, 3})
	assert.False(t, exists)
}
