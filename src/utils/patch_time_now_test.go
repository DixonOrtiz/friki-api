package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPatchTimeNow(t *testing.T) {
	PatchTimeNow()
	assert.Equal(t, time.Now(), time.Date(2023, 03, 06, 14, 27, 45, 3699162, time.UTC))
}
