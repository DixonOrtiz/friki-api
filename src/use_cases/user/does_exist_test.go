package userusecases

import (
	userrepo "frikiapi/src/adapters/repositories/user"
	"frikiapi/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesExist(t *testing.T) {
	var testUser = entities.User{
		ExternalID: "test_external_id",
		Name:       "Matías Ariel",
		LastName:   "Fernández Fernández",
		Email:      "matias.fernandez@gmail.com",
		Picture:    "https://www.latercera.com/resizer/-IUFynLZwc1L_2tw3OuYwnJErRk=/900x600/smart/cloudfront-us-east-1.images.arcpublishing.com/copesa/XP434CUEFNH7FMXNR4PAQ4RBAY.jpeg",
	}

	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByExternalID").Return(testUser, "test_user_doc", nil)

	userUseCases := MakeUserUseCases(userRepository)

	exist, document, err := userUseCases.DoesExist("test_external_id")

	assert.True(t, exist)
	assert.Equal(t, "test_user_doc", document)
	assert.Nil(t, err)
}
