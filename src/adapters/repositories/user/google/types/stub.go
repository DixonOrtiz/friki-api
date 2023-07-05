package types

import "frikiapi/src/entities"

var googleUserInput = GoogleUser{
	ID:         "12345",
	GivenName:  "Felipe Ignacio",
	FamilyName: "Flores Chandía",
	Email:      "felipe.flores@gmail.com",
	Picture:    "https://static.eldinamo.cl/media/2019/05/PS_06510.jpg",
}

var expectedUser = entities.User{
	Name:       "Felipe Ignacio",
	LastName:   "Flores Chandía",
	Email:      "felipe.flores@gmail.com",
	Picture:    "https://static.eldinamo.cl/media/2019/05/PS_06510.jpg",
	ExternalID: "12345",
}