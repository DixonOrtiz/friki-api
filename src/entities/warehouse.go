package entities

type Warehouse struct {
	ID        int    `json:"id"`
	StoreID   int    `json:"store_id"`
	AddressID int    `json:"address_id"`
	Name      string `json:"name"`
}
