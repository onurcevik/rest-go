package db

import "github.com/onurcevik/restful/internal/model"

func GetUserNotes(id int) ([]model.Note, error) {
	var notes []model.Note
	sqlstmnt := `SELECT id,note FROM notes WHERE notes.ownerid=$1`
	rows, err := Conn.Query(sqlstmnt, id)
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

func IsResourceOwner(resourceid, ownerid int) bool {
	var oid int
	sqlstmnt := `SELECT ownerid FROM notes  WHERE notes.id=$1`

	_ = Conn.QueryRow(sqlstmnt, resourceid, ownerid).Scan(&oid)

	if oid != ownerid {
		return false
	}
	return true
}
