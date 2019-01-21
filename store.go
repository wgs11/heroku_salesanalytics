package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type Store interface {
	CheckUser(creds *Credentials) error
	CreateUser(creds *NewUser) error
	GetUser(user_name string) (*NewUser, error)
	GetStore(user string) (*Location, error)
	GetReview(location string, day string) (*Review, error)
	GetReviews(location string, day string) ([]*Review, error)
	GetManagers() ([]*ManagerForm, error)
	CreateStore(creds *NewStoreCreds) error
	GetStores() ([]*Location, error)
	DBSeed(question string)
	GetQuestions() (questions []string)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) GetQuestions()(questions []string){
	rows, err := store.db.Query("SELECT * FROM questions")
	if err != nil {
		return nil
	} else {
		defer rows.Close()
		qs := []string{}
		var single string
		for rows.Next() {
			if err := rows.Scan(single); err != nil {
				return nil
			}
			qs = append(qs,single)
		}
		return qs
		fmt.Println(qs[0])
	}
	return nil
}

func (store *dbStore) DBSeed(question string) {
		row, err := store.db.Query("INSERT INTO questions (prompt) VALUES ($1)", question)
		if err != nil {
			fmt.Println("there was an issue with question creation.")
		}
		defer row.Close()
}

func (store *dbStore) GetUser(user_name string) (*NewUser, error) {
	user := &NewUser{}
	rows, err := store.db.Query("SELECT fname, lname, position, store_id FROM employees WHERE user_name = $1", user_name)
	if err != nil {
		fmt.Println("there was a problem")
		fmt.Println(err)
		return nil, err
	} else {
		defer rows.Close()
		err := rows.Next()
		if err {
			if err := rows.Scan(&user.First, &user.Last, &user.Position, &user.Home); err != nil {
				return nil, err
			}
		}
		return user, nil
	}
	return nil, nil
}

func (store *dbStore) CreateStore(creds *NewStoreCreds) error {
	_,err := store.db.Query("INSERT INTO stores (location_id,location_name,manager_id,region) VALUES ($1,$2,(SELECT employee_id FROM employees WHERE fname = $3),$4)", creds.Number, creds.Name, creds.First, creds.Region)
	if err != nil {
		fmt.Println("there was a problem")
		fmt.Println(err)
		return err
	}
	fmt.Println("got to here")
	return nil
}

func (store *dbStore) GetStores() ([]*Location, error) {
	rows, err := store.db.Query("SELECT location_name FROM stores")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	stores := []*Location{}
	for rows.Next() {
		single := &Location{}
		if err := rows.Scan(&single.City); err != nil {
			return nil, err
		}
		stores = append(stores, single)
	}
	return stores, nil
}

func (store *dbStore) GetManagers() ([]*ManagerForm, error) {
	rows, err := store.db.Query("SELECT employee_id, fname, lname FROM employees WHERE position > 3")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	managers := []*ManagerForm{}
	for rows.Next() {
		manager := &ManagerForm{}
		if err := rows.Scan(&manager.ID, &manager.First, &manager.Last); err != nil {
			return nil, err
		}
		managers = append(managers, manager)
	}
	return managers, nil
}

func (store *dbStore) GetReviews(location string, day string) ([]*Review, error) {
	rows, err := store.db.Query("SELECT day, answers::bit(100), feedback FROM reviews WHERE store_id = $1", location)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews := []*Review{}
	for rows.Next() {
		review := &Review{}
		review.Store_id, _ = strconv.Atoi(location)
		ans := []int8{}
		if err := rows.Scan(&review.Day, &ans, &review.Feedback); err != nil {
			return nil, err
		}
		review.Day = review.Day[:10]
		review.Outside = ans[:7]
		review.Emp_sys = ans[7:14]
		review.Eating = ans[14:30]
		review.Merch = ans[30:46]
		review.Fountain = ans[46:60]
		review.Inventory = ans[60:74]
		review.Backroom = ans[74:88]
		review.Restrooms = ans[88:100]
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (store *dbStore) GetReview(location string, date string) (*Review, error) {
	review := &Review{}
	row, err := store.db.Query("SELECT day, answers::bit(100), feedback FROM reviews WHERE day = $1 and store_id = $2", date, location)
	if err != nil {
		return nil, err
	} else {
		defer row.Close()
		err := row.Next()
		if err {
			ans := []int8{}
			if err := row.Scan(&review.Day, &ans, &review.Feedback); err != nil {
				return nil, err
			}
			review.Day = review.Day[:10]
			review.Outside = ans[:7]
			review.Emp_sys = ans[7:14]
			review.Eating = ans[14:30]
			review.Merch = ans[30:46]
			review.Fountain = ans[46:60]
			review.Inventory = ans[60:74]
			review.Backroom = ans[74:88]
			review.Restrooms = ans[88:100]
			return review, nil
		}
	}

	return nil, nil
}

func (store *dbStore) GetStore(user string) (*Location, error) {
	place := &Location{}
	row, err := store.db.Query("SELECT location_id,location_name,manager_id,region FROM stores WHERE location_id = (SELECT store_id FROM employees WHERE user_name = $1)",user)
	if err != nil {
		return nil, err
	} else {
		defer row.Close()
		err := row.Next()
		if err {
			row.Scan(&place.LocationID,&place.City,&place.ManagerID,&place.Region)
			fmt.Println(place.LocationID)
			return place, nil
		}
	}
	return nil, err
}
func (store *dbStore) CheckUser(creds *Credentials) error {
	dummyCreds := &Credentials{}
	row, err := store.db.Query("SELECT password FROM employees WHERE user_name = $1",creds.Username)
	if err != nil {
		return err
	} else {
		fmt.Println("for some reason there's no error here")
		defer row.Close()
		err := row.Next()
		if err {
			fmt.Println("for some reason it says there's a row")
			row.Scan(&dummyCreds.Password)
			fmt.Println(bcrypt.CompareHashAndPassword([]byte(dummyCreds.Password),[]byte(creds.Password)))
			return(bcrypt.CompareHashAndPassword([]byte(dummyCreds.Password),[]byte(creds.Password)))
		} else {
			return nil
		}
	}
	return nil
}

func (store *dbStore) CreateUser(creds *NewUser) error {
	fmt.Println(creds.Username, creds.Password, creds.First,creds.Last,creds.Position,creds.Home)
	hashedPassword,_ := bcrypt.GenerateFromPassword([]byte(creds.Password),8)
	fmt.Println(hashedPassword)
	fmt.Println("here we are")
	_,err := store.db.Query("INSERT INTO employees(fname,lname,position,store_id,user_name,password) VALUES ($1, $2, $3, (SELECT location_id FROM stores WHERE location_name = $4), $5, $6)", creds.First, creds.Last, creds.Position, creds.Home, string(creds.Username), string(hashedPassword))
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