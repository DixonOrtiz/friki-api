package types

type UpdateUserDTO struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Picture  string `json:"picture" binding:"required"`
}
