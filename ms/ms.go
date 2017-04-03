package dsds

import (
	"github.com/it512/dsds"
	"github.com/jmoiron/sqlx"
)

type (
	ReadOnlyer interface {
		IsReadOnly() bool
	}
)

type (
	MasterSlaveDbSet struct {
		Master *sqlx.DB
		Slave  dsds.DbSet
	}
)

func (db *MasterSlaveDbSet) GetDb(desc interface{}) (*sqlx.DB, error) {
	r, ok := desc.(ReadOnlyer)

	if ok || r.IsReadOnly() {
		return db.Slave.GetDb(desc)
	}

	return db.Master, nil
}
