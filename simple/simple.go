package simple

import "github.com/jmoiron/sqlx"

type (
	SimpleDbSet struct {
		*sqlx.DB
	}
)

func (db *SimpleDbSet) GetDb(desc interface{}) (*sqlx.DB, error) {
	return db.DB, nil
}

func NewSimpleDbSetWithDb(db *sqlx.DB) *SimpleDbSet {
	return &SimpleDbSet{DB: db}
}

func NewSimpleDbSet(driverName, dataSourceName string) *SimpleDbSet {
	db := sqlx.MustOpen(driverName, dataSourceName)
	return NewSimpleDbSetWithDb(db)
}
