package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
			fmt.Println(bcrypt.CompareHashAndPassword([]byte(dummyCreds.Password),[]byte(creds.Password)))
			return(bcrypt.CompareHashAndPassword([]byte(dummyCreds.Password),[]byte(creds.Password)))
		}
	}
	return nil
}

func (store *dbStore) CreateUser(creds *NewUser) error {
	fmt.Println(creds.Username, creds.Password, creds.First,creds.Last,creds.Position,creds.Home)
	hashedPassword,_ := bcrypt.GenerateFromPassword([]byte(creds.Password),8)
	fmt.Println(hashedPassword)
	fmt.Println("here we are")
	_,err := store.db.Query("INSERT INTO employees(fname,lname,position,store_id,user_name,password) VALUES ($1, $2, $3, $4, $5, $6)", creds.First, creds.Last, creds.Position, creds.Home, string(creds.Username), string(hashedPassword))
	if err != nil {
		fmt.Println("there was a problem")
		fmt.Println(err)
		return err
	}
	fmt.Println("got to here")
	return nil
}

var store Store

func InitStore(s Store) {
	store = s
}