package main

import (
	"log"
	"net/http"
)

func Displayhome(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "login", "")
	if err != nil {
		log.Fatal("Cannot retrieve login page.")
	}
}
