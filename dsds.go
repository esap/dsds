package dsds

import "github.com/jmoiron/sqlx"

type (
	DbSet interface {
		GetDb(desc interface{}) (*sqlx.DB, error)
	}
)
