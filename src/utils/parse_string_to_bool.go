package utils

import (
	"fmt"
	"strconv"
)

func ParseStringToBool(b string) (bool, error) {
	boolValue, err := strconv.ParseBool(b)
	if err != nil {
		return false, fmt.Errorf("entered value should be a boolean")
	}
	return boolValue, nil
}
