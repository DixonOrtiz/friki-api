package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStringToInt(t *testing.T) {
	number, err := ParseStringToInt("30")

	assert.Equal(t, 30, number)
	assert.Nil(t, err)
}

func TestParseNotNumberStringToInt(t *testing.T) {
	number, err := ParseStringToInt("not_number")

	assert.Equal(t, 0, number)
	assert.EqualError(t, err, "the string 'not_number' is not a number")
}
