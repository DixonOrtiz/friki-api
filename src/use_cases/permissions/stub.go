package permusecases

import "frikiapi/src/entities"

var (
	testUserID         = "test_user_id"
	testFirstAddressID = "test_first_address_id"
	testFirstStoreID   = "test_first_store_id"

	testPermission = entities.Permission{
		ID:     "test_permission_id",
		UserID: testUserID,
		Addresses: []string{
			testFirstAddressID,
		},
		Stores: []string{
			testFirstStoreID,
		},
	}
	emptyPermission = entities.Permission{
		ID:        "test_id",
		UserID:    "user_id",
		Addresses: []string{},
		Stores:    []string{},
	}
)
