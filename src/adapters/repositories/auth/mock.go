package authrepo

import (
	"github.com/stretchr/testify/mock"
)

type MockAuthRepository struct {
	mock.Mock
}

func (m MockAuthRepository) GenerateLoginURL() string {
	args := m.Called()
	result := args.Get(0)
	return result.(string)
}

func (m MockAuthRepository) GenerateExternalToken(code string) (string, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}
