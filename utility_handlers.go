package main

import (
	"fmt"
	"net/http"
)

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

type ManagerForm struct {
	ID int `json:"employee_id", db:"employee_id"`
	First string `json:"fname", db:"fname"`
	Last string `json:"lname", db:"lname"`
}

type NewUser struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
	First	string `json:"fname", db:"fname"`
	Last 	string `json:"fname", db:"lname"`
	Home	string `json:"store_id", db:"store_id"`
	Position string `json:"position", db:"position"`
}

func IsSignedIn(w http.ResponseWriter, r *http.Request) bool {
	session,_ := cache.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	} else {return true}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	creds := &Credentials{}
	creds.Username = r.FormValue("username")
	creds.Password = r.FormValue("password")
	if store.CheckUser(creds) == nil {
		session,_ := cache.Get(r, "cookie-name")
		session.Values["authenticated"] = true
		session.Values["user"] = creds.Username
		if (creds.Username == "Sheppy"){
			session.Values["admin"] = true
		} else {
			session.Values["admin"] = false
		}
		session.Save(r,w)
	} else{
	}
	http.Redirect(w,r,"/", 302)

}


func Signup(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	creds := &NewUser{}
	creds.Username = r.FormValue("username")
	creds.Password = r.FormValue("password")
	creds.First = r.FormValue("fname")
	creds.Last = r.FormValue("lname")
	creds.Position = r.FormValue("position")
	creds.Home = r.FormValue("store")
	fmt.Println(creds.Username, creds.Password, creds.First,creds.Last,creds.Position,creds.Home)
	store.CreateUser(creds)
	http.Redirect(w, r, "/", http.StatusOK)
}