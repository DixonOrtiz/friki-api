package input

type UpdateUserDTO struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Picture  string `json:"picture" binding:"required"`
}
