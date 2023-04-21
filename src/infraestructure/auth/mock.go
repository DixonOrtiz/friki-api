package authinfra

import (
	"context"

	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

type MockAuthConfig struct {
	mock.Mock
}

func (m MockAuthConfig) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	args := m.Called()
	result := args.Get(0)
	return result.(string)
}

func (m MockAuthConfig) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*oauth2.Token), args.Error(1)
}
