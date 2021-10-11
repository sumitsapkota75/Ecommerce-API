package models

// User -> Model

type User struct {
	UserBase
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserType string `json:"user_type,omitempty"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

// TableName  -> returns table name of model
func (u User) TableName() string {
	return "users"
}
