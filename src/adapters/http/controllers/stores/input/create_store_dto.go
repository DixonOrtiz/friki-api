package input

type CreateStoreDTO struct {
	ID          string `json:"wid"`
	UserID      string `json:"user_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
