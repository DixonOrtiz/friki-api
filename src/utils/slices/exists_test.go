package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExistsInSlice(t *testing.T) {
	list := []string{"test_store_id_1", "test_store_id_2"}
	exists := Exists(list, "test_store_id_2")
	assert.True(t, exists)
}

func TestDoesNotExistsInSlice(t *testing.T) {
	list := []string{"test_store_id_1", "test_store_id_2"}
	exists := Exists(list, "test_store_id_3")
	assert.False(t, exists)
}
