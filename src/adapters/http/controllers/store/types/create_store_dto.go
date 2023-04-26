package types

type CreateStoreDTO struct {
	Name             string     `json:"name" binding:"required"`
	Description      string     `json:"description" binding:"required"`
	Address          AddressDTO `json:"address"`
	WarehouseAddress AddressDTO `json:"warehouse_address"`
}

type AddressDTO struct {
	Name           string `json:"name"`
	Region         string `json:"region" binding:"required"`
	City           string `json:"city" binding:"required"`
	Commune        string `json:"commune" binding:"required"`
	Street         string `json:"street" binding:"required"`
	Number         string `json:"number" binding:"required"`
	AdditionalInfo string `json:"additional_info"`
}
