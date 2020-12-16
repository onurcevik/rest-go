package model

//Note type
type Note struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

func (n *Note) TableName() string {
	return "notes"
}
