package entities

type User struct {
	ID         int    `json:"id"`
	RoleID     int    `json:"role_id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Cellphone  string `json:"cellphone"`
	Rut        string `json:"rut"`
	ExternalID string `json:"external_id"`
	Picture    string `json:"picture"`
}
