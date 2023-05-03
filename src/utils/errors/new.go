package errors

import (
	"fmt"
)

func New(code string, err interface{}) error {
	message := ""

	switch e := err.(type) {
	case string:
		message = e
	case error:
		message = e.Error()
	}

	return fmt.Errorf("%s: %s", code, message)
}
