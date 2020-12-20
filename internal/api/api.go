package api

import (
	"github.com/onurcevik/rest-go/internal/db"
)

type API struct {
	Database *db.Database
}

func (a *API) GetDB() *db.Database {
	return a.Database
}
