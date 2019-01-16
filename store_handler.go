package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)
type Location struct {
	LocationID int `json:"location_id", db:"location_id"`
	City string `json:"location_name", db:"location_name"`
	ManagerID	int `json:"manager_id", db:"manager_id"`
	Region 	string `json:"region", db:"region"`
}

type NewStoreCreds struct {
	Number int `json:"location_id", db:"location_id"`
	First string `json:"fname", db:"fname"`
	Last string `json:"lname", db:"lname"`
	Name string `json:"location_name", db:"location_name"`
	Region string `json:"region", db:"region"`
}

func Getstores(w http.ResponseWriter, r *http.Request) {
	session,_ := cache.Get(r, "cookie-name")
	if user, ok := session.Values["user"].(string); ok {
		place, err := store.GetStore(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			err = templates.ExecuteTemplate(w, "store", (*place))
			if err != nil {
				log.Fatal("Cannot retrieve store page.")
			}
		}
	} else {
		fmt.Println("what just happened")
	}



}

func Makestore(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	new := &NewStoreCreds{}
	new.Number,_ = strconv.Atoi(r.FormValue("storeid"))
	manager := r.FormValue("manager")
	names := strings.Fields(manager)
	new.First = names[0]
	new.Last = names[1]
	new.Name = r.FormValue("location")
	new.Region = r.FormValue("region")
	err := store.CreateStore(new)
	if err != nil {
		http.Redirect(w,r,"/",404)
	}
	http.Redirect(w,r,"/",302)
}


