package main

import (
	"fmt"
	"log"
	"net/http"
)


func Displaycreate(w http.ResponseWriter, r *http.Request) {
	stores, err := store.GetStores()
	storeblock := []Location{}
	for _, element := range stores {
		storeblock = append(storeblock, (*element))
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	err = templates.ExecuteTemplate(w, "create", stores)
	if err != nil {
		log.Fatal("Cannot retrieve login page.")
	}
}

func Displaycreatestore(w http.ResponseWriter, r *http.Request) {
	managers, err := store.GetManagers()
	if err != nil {
		fmt.Println(err)
		return
	}
	managerblock := []ManagerForm{}
	for _, element := range managers {
		managerblock = append(managerblock, (*element))
	}
	err = templates.ExecuteTemplate(w, "create_store", managerblock)

}
func Displaystores(w http.ResponseWriter, r *http.Request) {

}

func Newreview(w http.ResponseWriter, r *http.Request) {
	qs := store.GetQuestions()
	fmt.Println(qs[0])
	err := templates.ExecuteTemplate(w, "newreview", qs)
	if err != nil {
		log.Fatal("Cannot retrieve new review page.")
	}
}

func Displayprofile(w http.ResponseWriter, r *http.Request) {
	//session,_ := cache.Get(r, "cookie-name")
	//if str, ok := session.Values["user"].(string); ok {
	//user,_ := store.GetUser(str)
	//if user != nil {
	//err = templates.ExecuteTemplate(w, "profile", user)
	//}
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

