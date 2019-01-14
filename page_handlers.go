package main

import (
	"log"
	"net/http"
)

func Makeuser(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "create", "")
	if err != nil {
		log.Fatal("Cannot retrieve login page.")
	}
}

func Displayhome(w http.ResponseWriter, r *http.Request) {
	var err error
	if IsSignedIn(w,r) {
		err = templates.ExecuteTemplate(w, "home", "")
	} else {
		err = templates.ExecuteTemplate(w, "login", "")
	}
	if err != nil {
		log.Fatal("Cannot retrieve page.")
	}
}
