package permrepo

import (
	"frikiapi/src/entities"

	"github.com/stretchr/testify/mock"
)

type MockPermissionRepository struct {
	mock.Mock
}

func (m *MockPermissionRepository) Create(permission entities.Permission) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockPermissionRepository) AddResource(document string, permission entities.Permission) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockPermissionRepository) GetByUserID(userID string) (entities.Permission, string, error) {
	args := m.Called()
	return args.Get(0).(entities.Permission), args.Get(1).(string), args.Error(1)
}
