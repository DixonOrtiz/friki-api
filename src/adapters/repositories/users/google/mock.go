package googleuserrepo

import (
	"frikiapi/src/entities"

	"github.com/stretchr/testify/mock"
)

type MockGoogleUserRepository struct {
	mock.Mock
}

func (m *MockGoogleUserRepository) GetByToken(token string) (entities.User, error) {
	args := m.Called()
	return args.Get(0).(entities.User), args.Error(1)
}
