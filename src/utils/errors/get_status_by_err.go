package errors

import (
	"net/http"
	"strings"
)

var errorsAndStatuses = map[string]int{
	INTERNAL:      http.StatusInternalServerError,
	UNAUTHORIZED:  http.StatusUnauthorized,
	UNPROCESSABLE: http.StatusUnprocessableEntity,
	NOT_FOUND:     http.StatusNotFound,
	BAD_REQUEST:   http.StatusBadRequest,
}

func GetStatusByErr(err error) int {
	for internalErr, status := range errorsAndStatuses {
		if strings.Contains(err.Error(), internalErr) {
			return status
		}
	}
	return http.StatusInternalServerError
}
