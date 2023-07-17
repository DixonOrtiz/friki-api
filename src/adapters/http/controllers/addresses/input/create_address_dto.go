package input

type CreateAddressDTO struct {
	Name             string `json:"name" binding:"required"`
	City             string `json:"city" binding:"required"`
	Commune          string `json:"commune" binding:"required"`
	Street           string `json:"street" binding:"required"`
	Number           int    `json:"number" binding:"required"`
	Apartment        string `json:"apartment" binding:"required"`
	Sector           string `json:"sector" binding:"required"`
	Type             string `json:"type" binding:"required"`
	ExtraInformation string `json:"extra_information" binding:"required"`
}
