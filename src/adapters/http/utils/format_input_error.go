package utils

import (
	"strings"
)

func FormatInputError(err error) []string {
	return strings.Split(err.Error(), "\n")
}
