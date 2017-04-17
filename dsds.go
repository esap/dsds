package dsds

import "github.com/jmoiron/sqlx"

type (
	DbManager interface {
		GetDb(desc interface{}) (*sqlx.DB, error)
	}
)

func MustGetDb(db *sqlx.DB, e error) *sqlx.DB {
	if e != nil {
		panic(e)
	}
	return db
}
