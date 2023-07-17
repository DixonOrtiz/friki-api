package permusecases

import "frikiapi/src/entities"

var (
	testUserID         = "test_user_id"
	testFirstAddressID = "test_first_address_id"
	testPermission     = entities.Permission{
		ID:     "test_permission_id",
		UserID: testUserID,
		Addresses: []string{
			testFirstAddressID,
		},
	}
)
