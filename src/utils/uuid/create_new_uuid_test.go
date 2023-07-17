package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	result := New()
	assert.NotEmpty(t, result)
}
