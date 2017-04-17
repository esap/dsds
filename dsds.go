package dsds

import "github.com/jmoiron/sqlx"

type (
	DbManager interface {
		GetDb(desc interface{}) (*sqlx.DB, error)
	}
)

func MustGetDb(dbm DbManager, i interface{}) *sqlx.DB {
	return Must(dbm.GetDb(i))
}

func Must(db *sqlx.DB, e error) *sqlx.DB {
	if e != nil {
		panic(e)
	}
	return db
}
