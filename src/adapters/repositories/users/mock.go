package userrepo

import (
	"frikiapi/src/entities"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByID(ID string) (entities.User, string, error) {
	args := m.Called()
	user := args.Get(0)
	document := args.Get(1)
	return user.(entities.User), document.(string), args.Error(2)
}

func (m *MockUserRepository) GetByExternalID(externalID string) (entities.User, string, error) {
	args := m.Called()
	user := args.Get(0)
	document := args.Get(1)
	return user.(entities.User), document.(string), args.Error(2)
}

func (m *MockUserRepository) Create(user entities.User) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserRepository) Update(document string, user entities.User) error {
	args := m.Called()
	return args.Error(0)
}
