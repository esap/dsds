package dsds

import "github.com/jmoiron/sqlx"

type (
	SimpleDbSet struct {
		*sqlx.DB
	}
)

func (db *SimpleDbSet) GetDb(desc interface{}) (*sqlx.DB, error) {
	return db.DB, nil
}

func NewSimpleDbSet(db *sqlx.DB) *SimpleDbSet {
	return &SimpleDbSet{DB: db}
}
