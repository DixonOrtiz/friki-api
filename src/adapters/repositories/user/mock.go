package userrepo

import (
	"frikiapi/src/entities"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m MockUserRepository) GetByExternalID(externalID string) (entities.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(entities.User), args.Error(1)
}

func (m MockUserRepository) GetByID(userID int) (entities.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(entities.User), args.Error(1)
}

func (m MockUserRepository) GetExternalUserByToken(token string) (entities.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(entities.User), args.Error(1)
}

func (m MockUserRepository) Register(user entities.User) error {
	args := m.Called()
	return args.Error(0)
}
