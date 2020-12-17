package api

import (
	"github.com/gobuffalo/envy"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/onurcevik/restful/internal/db"
)

var (
	suitapi *API
)

type Suite struct {
	API API
}

func NewSuite() *Suite {
	appWD, _ := os.Getwd()
	appWD = strings.Split(appWD, "/rest-go")[0] + "/rest-go"

	if err := envy.Load(filepath.Join(appWD, ".env.test")); err != nil {
		log.Panic(err)
	}

	config := &db.PostgresConfig{
		Host:     os.Getenv("pghost"),
		Port:     5432,
		User:     os.Getenv("pguser"),
		Password: os.Getenv("pgpass"),
		DBname:   os.Getenv("pgdbname"),
	}

	d := db.NewDB(*config)
	suitapi = new(API)
	suitapi.Database = &d

	return &Suite{API: *suitapi}

}
