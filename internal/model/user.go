package model

//User type
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "users"
}
