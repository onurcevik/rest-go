package db

import (
	"database/sql"

	"github.com/onurcevik/restful/internal/model"
)

var (
	Conn *sql.DB
)

func GetUserNotes(username string) []model.Note {
	var notes []model.Note
	sqlstmnt := `SELECT id,note FROM notes WHERE notes.username=$1`
	rows, err := Conn.Query(sqlstmnt, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var content string

		err = rows.Scan(&id, &content)
		if err != nil {
			panic(err)
		}
		notes = append(notes, model.Note{ID: id, Content: content})

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return notes
}
