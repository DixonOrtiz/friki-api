package errors

import (
	"errors"
	consterrors "frikiapi/src/utils/consts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	inputError := errors.New("there was an internal arror")
	err := New(consterrors.INTERNAL, inputError)
	assert.EqualError(t, err, "internal: there was an internal arror")
}

func TestNewStringError(t *testing.T) {
	inputError := "the resource was not found"
	err := New(consterrors.NOT_FOUND, inputError)
	assert.EqualError(t, err, "not_found: the resource was not found")
}
