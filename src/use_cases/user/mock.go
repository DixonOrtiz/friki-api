package userusecases

import (
	"frikiapi/src/entities"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCases struct {
	mock.Mock
}

func (m *MockUserUseCases) Create(user entities.User) (bool, error) {
	args := m.Called()
	return args.Get(0).(bool), args.Error(1)
}

func (m *MockUserUseCases) Update(user entities.User) error {
	args := m.Called()
	return args.Error(0)
}
