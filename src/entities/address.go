package entities

type Address struct {
	ID             int     `json:"id"`
	StoreID        int     `json:"store_id"`
	Region         string  `json:"region"`
	City           string  `json:"city"`
	Commune        string  `json:"commune"`
	Street         string  `json:"street"`
	Number         string  `json:"number"`
	AdditionalInfo *string `json:"additional_info"`
}
