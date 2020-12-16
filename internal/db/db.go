package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

//Database struct is used to transfer database connection between modules
type Database struct {
	Conn *sql.DB
}

//DBResource is usedfor  creating interfacee between models and databese methods
type DBResource interface {
	TableName() string
}

func getPostgresDataSource() string {

	host := os.Getenv("pghost")
	port, err := strconv.Atoi(os.Getenv("pgport"))
	if err != nil {
		log.Fatalln()
	}
	user := os.Getenv("pguser")
	password := os.Getenv("pgpass")
	dbname := os.Getenv("pgdbname")

	return fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

//Creates and returns new Database with connection
func NewDB() Database {
	d := Database{}
	psqlInfo := getPostgresDataSource()
	var err error
	d.Conn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = d.Conn.Ping()
	if err != nil {
		panic(err)
	}
	return d
}

func (db *Database) Insert(dbr DBResource, parameters ...interface{}) (int, error) {
	//return values
	var insertid int

	var stmnt string
	switch table := dbr.TableName(); table {
	case "users":
		stmnt = `INSERT INTO users (username, password)
		VALUES ($1,$2 ) RETURNING id;`
	case "notes":
		stmnt = `INSERT INTO notes  (ownerid, content)
		VALUES ($1,$2 ) RETURNING id;`
	default:
		return -1, fmt.Errorf("Unkown interface type")

	}

	err := db.Conn.QueryRow(stmnt, parameters[0], parameters[1]).Scan(&insertid)
	if err != nil {
		return -1, nil
	}
	return insertid, nil
}

func (db *Database) Update(dbr DBResource, parameters ...interface{}) (int, error) {
	//return values
	var insertid int
	var stmnt string
	switch table := dbr.TableName(); table {
	case "users":
		//TODO implement
	case "notes":
		stmnt = `UPDATE notes SET content=$2 where id=$1 RETURNING id;`
	default:
		return -1, fmt.Errorf("Unkown interface type")

	}

	err := db.Conn.QueryRow(stmnt, parameters[0], parameters[1]).Scan(&insertid)
	if err != nil {
		return -1, nil
	}
	return insertid, nil

}
func (db *Database) Delete(dbr DBResource, id int) (int, error) {
	//return values
	var stmnt string
	var insertid int
	switch table := dbr.TableName(); table {
	case "users":
		stmnt = `DELETE  FROM users WHERE id=$1 RETURNING id;`
	case "notes":
		stmnt = `DELETE  FROM notes WHERE id=$1 RETURNING id;`
	default:
		return -1, fmt.Errorf("Unkown interface type")

	}
	err := db.Conn.QueryRow(stmnt, id).Scan(&insertid)
	if err != nil {
		return -1, err
	}

	return insertid, nil
}
