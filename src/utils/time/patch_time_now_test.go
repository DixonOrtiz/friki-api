package time

import (
	"testing"
	gotime "time"

	"github.com/stretchr/testify/assert"
)

func TestPatchTimeNow(t *testing.T) {
	PatchTimeNow()
	assert.Equal(t, gotime.Now(), gotime.Date(2023, 03, 06, 14, 27, 45, 3699162, gotime.UTC))
}
