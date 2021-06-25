package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

func (user User) TableName() string {
	return "users"
}
