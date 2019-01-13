package main


import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Store interface {
	CheckUser(creds *Credentials) error
	CreateUser(creds *NewUser) error

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

func (store *dbStore) CreateUser(creds *NewUser) error {
	hashedPassword,_ := bcrypt.GenerateFromPassword([]byte(creds.Password),8)
	_,err := store.db.Query("INSERT INTO employees(fname,lname,position,store_id,user_name,password) VALUES ($1, $2, $3, $4, $5, $6)", string(creds.First), string(creds.Last), creds.Position, 3, string(creds.Username), string(hashedPassword))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

var store Store

func InitStore(s Store) {
	store = s
}