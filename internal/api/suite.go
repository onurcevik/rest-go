package api

import (
	"os"

	"github.com/onurcevik/restful/internal/db"
)

var (
	suitapi *API
)

type Suite struct {
	API API
}

func NewSuite() *Suite {

	config := &db.PostgresConfig{
		Host:     os.Getenv("pghost"),
		Port:     5432,
		User:     os.Getenv("pguser"),
		Password: os.Getenv("pgpass"),
		DBname:   os.Getenv("testpgdbname"),
	}

	d := db.NewDB(*config)
	suitapi.Database = &d

	return &Suite{API: *suitapi}

}
