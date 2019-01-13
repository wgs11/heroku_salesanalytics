package main

import (
  "database/sql"
  "github.com/gorilla/mux"
  "log"
  "fmt"
  "net/http"
  "os"
  "html/template"
  "github.com/gorilla/securecookie"
  "github.com/gorilla/sessions"
)
var fmap = template.FuncMap {

}
var templates = template.Must(template.New("").Funcs(fmap).ParseGlob("templates/*gohtml"))
var cache = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

func newRouter() *mux.Router {
  r := mux.NewRouter()
  staticFileDirectory := http.Dir("./assets/")
  staticFileHandler := http.StripPrefix("/assets/",http.FileServer(staticFileDirectory))
  r.HandleFunc("/signin", Signin)
  r.HandleFunc("/", Displayhome)
  r.PathPrefix("/assets/").Handler(staticFileHandler)
  r.HandleFunc("/stores", Getstores)
  r.HandleFunc("/create", Makeuser)
  r.HandleFunc("/signup", Signup)
  return r
}

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello World")
}

func main() {
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  InitStore(&dbStore{db: db})
  db.Ping()
  r := newRouter()
  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, r); err != nil {
    panic(err)
  }
}