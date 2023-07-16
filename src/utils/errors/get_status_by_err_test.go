package errors

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInternalStatusByError(t *testing.T) {
	status := GetStatusByErr(fmt.Errorf("%s: internal error", INTERNAL))
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestGetUnauthorizedStatusByError(t *testing.T) {
	status := GetStatusByErr(fmt.Errorf("%s: the user does not have access", UNAUTHORIZED))
	assert.Equal(t, http.StatusUnauthorized, status)
}

func TestGetConflictStatusByError(t *testing.T) {
	status := GetStatusByErr(fmt.Errorf("%s: there is no data to process", CONFLICT))
	assert.Equal(t, http.StatusConflict, status)
}

func TestGetNotFoundStatusByError(t *testing.T) {
	status := GetStatusByErr(fmt.Errorf("%s: the user does not exist", NOT_FOUND))
	assert.Equal(t, http.StatusNotFound, status)
}

func TestGetBadRequestStatusByError(t *testing.T) {
	status := GetStatusByErr(fmt.Errorf("%s: user_id param is required", BAD_REQUEST))
	assert.Equal(t, http.StatusBadRequest, status)
}
