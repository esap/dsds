package dsds

import "github.com/jmoiron/sqlx"

type (
	DbManager interface {
		GetDb(desc interface{}) (*sqlx.DB, error)
	}
)

func MustGetDb(dbm *DbManager, i interface{}) *sqlx.DB {
	db, e := dbm.GetDb(i)
	if e != nil {
		panic(e)
	}
	return db
}
