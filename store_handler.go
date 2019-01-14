package main

import (
	"fmt"
	"net/http"
)
type Location struct {
	LocationID int `json:"location_id", db:"location_id"`
	City string `json:"location_name", db:"location_name"`
	ManagerID	int `json:"manager_id", db:"manager_id"`
	Region 	string `json:"region", db:"region"`
}

func Getstores(w http.ResponseWriter, r *http.Request) {
	session,_ := cache.Get(r, "cookie-name")
	if user, ok := session.Values["user"].(string); ok {
		store.GetStore(user)
	} else {
		fmt.Println("what just happened")
	}



}


