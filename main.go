package main

import (
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
  r.PathPrefix("/assets/").Handler(staticFileHandler)
  r.HandleFunc("/", Displayhome)
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
  r := newRouter()

  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, r); err != nil {
    panic(err)
  }
}