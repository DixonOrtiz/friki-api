package types

import (
	"fmt"
	"strconv"
)

func ParseStringToInt(str string) (int, error) {
	number, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("the string '%s' is not a number", str)
	}

	return number, nil
}
