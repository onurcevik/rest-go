package db

import (
	"fmt"

	"github.com/onurcevik/restful/internal/model"
)

func (db Database) DoesUserExist(username string) bool {
	var usernameExists bool
	sqlstmnt := `SELECT EXISTS(SELECT * FROM users WHERE username=$1);`
	_ = db.Conn.QueryRow(sqlstmnt, username).Scan(&usernameExists)
	return usernameExists
}

func (db Database) GetUserNotes(id int) ([]model.Note, error) {
	var notes []model.Note
	sqlstmnt := `SELECT id,note FROM notes WHERE notes.ownerid=$1`
	rows, err := db.Conn.Query(sqlstmnt, id)
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

func (db Database) IsResourceOwner(resourceid, ownerid int) bool {
	var oid int
	sqlstmnt := `SELECT ownerid FROM notes  WHERE notes.id=$1`
	_ = db.Conn.QueryRow(sqlstmnt, resourceid).Scan(&oid)
	fmt.Println("oid", oid)
	fmt.Println("ownerid", ownerid)
	if oid != ownerid {
		return false
	}
	return true
}
