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

func (m *MockAddressRepository) GetByID(ID string) (entities.Address, string, error) {
	args := m.Called()
	return args.Get(0).(entities.Address), args.Get(1).(string), args.Error(2)
}

func (m *MockAddressRepository) Update(document string, address entities.Address) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockAddressRepository) Delete(document string) error {
	args := m.Called()
	return args.Error(0)
}
