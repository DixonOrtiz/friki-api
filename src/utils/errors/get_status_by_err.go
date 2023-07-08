package errors

import (
	"frikiapi/src/utils/consts"
	"net/http"
	"strings"
)

var errorsAndStatuses = map[string]int{
	consts.Errors.INTERNAL:      http.StatusInternalServerError,
	consts.Errors.UNAUTHORIZED:  http.StatusUnauthorized,
	consts.Errors.UNPROCESSABLE: http.StatusUnprocessableEntity,
	consts.Errors.NOT_FOUND:     http.StatusNotFound,
	consts.Errors.BAD_REQUEST:   http.StatusBadRequest,
}

func GetStatusByErr(err error) int {
	for internalErr, status := range errorsAndStatuses {
		if strings.Contains(err.Error(), internalErr) {
			return status
		}
	}
	return http.StatusInternalServerError
}
