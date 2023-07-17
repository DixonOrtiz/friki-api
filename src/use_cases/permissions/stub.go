package permusecases

import "frikiapi/src/entities"

var testPermission = entities.Permission{
	ID:     "test_permission_id",
	UserID: "test_user_id",
	Addresses: []string{
		"test_first_address_id",
	},
}
