package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	result := RemoveElement([]string{"apple", "banana", "orange"}, "banana")
	assert.Equal(t, []string{"apple", "orange"}, result)
}

func TestRemoveElementThatIsDuplicated(t *testing.T) {
	result := RemoveElement([]string{"apple", "banana", "banana", "orange"}, "banana")
	assert.Equal(t, []string{"apple", "banana", "orange"}, result)
}

func TestRemoveElementTwice(t *testing.T) {
	result := RemoveElement([]string{"apple", "banana", "banana", "orange"}, "banana")
	result = RemoveElement(result, "banana")
	assert.Equal(t, []string{"apple", "orange"}, result)
}

func TestRemoveThatDoesNotExist(t *testing.T) {
	result := RemoveElement([]string{"apple", "banana", "orange"}, "grape")
	assert.Equal(t, []string{"apple", "banana", "orange"}, result)
}

func TestRemoveFirstElement(t *testing.T) {
	result := RemoveElement([]string{"apple", "banana", "orange"}, "apple")
	assert.Equal(t, []string{"banana", "orange"}, result)
}

func TestRemoveLastElement(t *testing.T) {
	result := RemoveElement([]string{"apple", "banana", "orange"}, "orange")
	assert.Equal(t, []string{"apple", "banana"}, result)
}
