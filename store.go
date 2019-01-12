package main


import (
	"database/sql"
)

type Store interface {

}

type dbStore struct {
	db *sql.DB
}


var store Store

func InitStore(s Store) {
	store = s
}