package permusecases

import (
	"frikiapi/src/entities"

	"github.com/stretchr/testify/mock"
)

type MockPermissionUseCases struct {
	mock.Mock
}

func (m *MockPermissionUseCases) Create(userID string) (entities.Permission, error) {
	args := m.Called()
	return args.Get(0).(entities.Permission), args.Error(1)
}

func (m *MockPermissionUseCases) AddResource(resource string, userID string, resourceID string) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockPermissionUseCases) Authorize(JWTUserID string, userID string, resources map[string]string) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockPermissionUseCases) RemoveResource(resource string, userID string, resourceID string) error {
	args := m.Called()
	return args.Error(0)
}
