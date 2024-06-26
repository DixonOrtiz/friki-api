package userusecases

import (
	"frikiapi/src/entities"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCases struct {
	mock.Mock
}

func (m *MockUserUseCases) Create(user entities.User) (string, bool, error) {
	args := m.Called()
	return args.Get(0).(string), args.Get(1).(bool), args.Error(2)
}

func (m *MockUserUseCases) Update(user entities.User) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserUseCases) GetByID(iD string) (entities.User, error) {
	args := m.Called()
	return args.Get(0).(entities.User), args.Error(1)
}
