package dsds

import "github.com/jmoiron/sqlx"

type (
	DbManager interface {
		GetDb(desc interface{}) (*sqlx.DB, error)
	}
)
