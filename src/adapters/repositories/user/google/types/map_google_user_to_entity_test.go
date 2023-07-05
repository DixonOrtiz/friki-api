package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapGoogleUserToEntity(t *testing.T) {
	user := MapGoogleUserToEntity(googleUserInput)
	assert.Equal(t, expectedUser, user)
}