package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatDate(t *testing.T) {
	PatchTimeNow()
	date := time.Now()
	formattedDate := FormatDate(date)
	assert.Equal(t, "2023-03-06 11:27:45", formattedDate)
}
