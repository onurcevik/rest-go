package model

//User type
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Note type
type Note struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}
