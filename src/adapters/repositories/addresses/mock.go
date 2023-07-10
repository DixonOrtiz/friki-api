package addressrepo

import (
	"frikiapi/src/entities"

	"github.com/stretchr/testify/mock"
)

type MockAddressRepository struct {
	mock.Mock
}

func (m *MockAddressRepository) Create(address entities.Address) error {
	args := m.Called()
	return args.Error(0)
}
