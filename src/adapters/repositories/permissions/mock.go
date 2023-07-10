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
