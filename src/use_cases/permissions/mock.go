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
