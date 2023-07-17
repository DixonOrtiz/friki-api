package addressusecases

import (
	"frikiapi/src/entities"
)

func (u AddressUseCases) Update() (entities.Address, error) {
	return entities.Address{}, nil
	// foundUser, document, err := u.AddressRepository.GetByID()
	// if err != nil {
	// 	return errors.New(errors.INTERNAL, err)
	// }

	// if foundUser.ID == "" {
	// 	return errors.New(errors.NOT_FOUND, fmt.Sprintf(
	// 		"user with id '%s' is not in the registers",
	// 		user.ID,
	// 	))
	// }

	// user.ExternalID = foundUser.ExternalID
	// user.Email = foundUser.Email
	// user.UpdatedAt = time.Now()
	// user.CreatedAt = foundUser.CreatedAt

	// err = u.UserRepository.Update(document, user)
	// if err != nil {
	// 	return errors.New(errors.INTERNAL, err)
	// }

	// return nil
}
