package db

import "github.com/onurcevik/restful/internal/model"

func GetUserNotes(username string) ([]model.Note, error) {
	var notes []model.Note
	sqlstmnt := `SELECT id,note FROM notes WHERE notes.username=$1`
	rows, err := Conn.Query(sqlstmnt, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var content string

		err = rows.Scan(&id, &content)
		if err != nil {
			return nil, err
		}
		notes = append(notes, model.Note{ID: id, Content: content})

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return notes, nil
}
