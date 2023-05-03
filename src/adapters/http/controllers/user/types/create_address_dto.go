package types

type CreateAddressDTO struct {
	Region         string  `json:"region" binding:"required"`
	City           string  `json:"city" binding:"required"`
	Commune        string  `json:"commune" binding:"required"`
	Street         string  `json:"street" binding:"required"`
	Number         string  `json:"number" binding:"required"`
	AdditionalInfo *string `json:"additional_info"`
}
