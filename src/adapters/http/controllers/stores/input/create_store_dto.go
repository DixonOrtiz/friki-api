package input

type CreateStoreDTO struct {
	ID          string `json:"wid"`
	UserID      string `json:"user_id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
