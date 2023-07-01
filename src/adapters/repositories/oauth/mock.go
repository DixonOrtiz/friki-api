package oauthrepo

import (
	"github.com/stretchr/testify/mock"
)

type MockOAuthRepository struct {
	mock.Mock
}

func (m *MockOAuthRepository) GenerateLoginURL() string {
	args := m.Called()
	result := args.Get(0)
	return result.(string)
}

func (m *MockOAuthRepository) GenerateExternalToken(code string) (string, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}
