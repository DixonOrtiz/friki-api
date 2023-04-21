package types

import "frikiapi/src/entities"

func MapGoogleUserToEntity(googleUser GoogleUser) entities.User {
	return entities.User{
		Name:       googleUser.GivenName,
		LastName:   googleUser.FamilyName,
		Email:      googleUser.Email,
		Picture:    googleUser.Picture,
		ExternalID: googleUser.ID,
	}
}
