package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
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
		fmt.Println("/")
	}

}


func Signup(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(creds.Username, creds.Password)
	store.CreateUser(creds)
}