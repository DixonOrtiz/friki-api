package errors

import (
	"frikiapi/src/utils/consts"
	"net/http"
	"strings"
)

var errorsAndStatuses = map[string]int{
	consts.INTERNAL:      http.StatusInternalServerError,
	consts.UNAUTHORIZED:  http.StatusUnauthorized,
	consts.UNPROCESSABLE: http.StatusUnprocessableEntity,
	consts.NOT_FOUND:     http.StatusNotFound,
	consts.BAD_REQUEST:   http.StatusBadRequest,
}

func GetStatusByErr(err error) int {
	for internalErr, status := range errorsAndStatuses {
		if strings.Contains(err.Error(), internalErr) {
			return status
		}
	}
	return http.StatusInternalServerError
}
