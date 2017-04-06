package mulit

import (
	"math/rand"

	"github.com/jmoiron/sqlx"
)

type (
	SelectFunc func(i interface{}, keys []string) (string, error)

	MulitDbSet struct {
		DbMap      map[string]*sqlx.DB
		SelectFunc SelectFunc
		Keys       []string
	}
)

func (db *MulitDbSet) GetDb(desc interface{}) (*sqlx.DB, error) {
	key, e := db.SelectFunc(desc, db.Keys)
	if e != nil {
		return nil, e
	}
	return db.DbMap[key], nil
}

func RandomKey(i interface{}, keys []string) (string, error) {
	return keys[rand.Intn(len(keys))], nil
}
