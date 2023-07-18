package addressusecases

import "frikiapi/src/entities"

var testAddressID = "test_address_id"

var testAddress = entities.Address{
	ID:               "test_address_id",
	UserID:           "test_user_id",
	Name:             "Mi casa",
	City:             "Santiago",
	Commune:          "Santiago Centro",
	Street:           "Compañía de Jesús",
	Number:           261,
	Apartment:        "0209",
	Sector:           "A",
	Type:             "home",
	ExtraInformation: "Dejas en conserjería.",
}
