package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatInputError(t *testing.T) {
	testError := errors.New("Key: 'RequestRoleDTO.Role' Error:Field validation for 'Role' failed on the 'required' tag\nKey: 'RequestRoleDTO.Reason' Error:Field validation for 'Reason' failed on the 'required' tag")
	expectedError := []string{
		"Key: 'RequestRoleDTO.Role' Error:Field validation for 'Role' failed on the 'required' tag",
		"Key: 'RequestRoleDTO.Reason' Error:Field validation for 'Reason' failed on the 'required' tag",
	}
	assert.Equal(t, expectedError, FormatInputError(testError))
}
