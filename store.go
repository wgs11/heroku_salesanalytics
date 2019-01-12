package main


import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type Store interface {
	CheckUser(creds *Credentials) error
	CreateUser(creds *Credentials) error

}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CheckUser(creds *Credentials) error {
	dummyCreds := &Credentials{}
	row, err := store.db.Query("SELECT password FROM users WHERE username = $1",creds.Username)
	if err != nil {
		return err
	} else {
		defer row.Close()
		err := row.Next()
		if err {
			row.Scan(&dummyCreds.Password)
			return(bcrypt.CompareHashAndPassword([]byte(dummyCreds.Password),[]byte(creds.Password)))
		}
	}
	return nil
}

func (store *dbStore) CreateUser(creds *Credentials) error {
	hashedPassword,_ := bcrypt.GenerateFromPassword([]byte(creds.Password),8)
	_,err := store.db.Query("INSERT INTO employees VALUES ($1, $2)", string(creds.Username), string(hashedPassword))
	if err != nil {
		return err
	}
	return nil
}

var store Store

func InitStore(s Store) {
	store = s
}