package main

import (
	"VinTravel/routes"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
)

const (
  host = "localhost"
  port = "5432"
  user = "postgres"
  password = "1234"
  dbname = "postgres"
)

type register_data struct {
  Username string `json:"username"`
  User_email string `json:"user_email"`
  Password string `json:"password"`
  Name string `json:"name"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "heelo")
}

func test_post(w http.ResponseWriter, r *http.Request) {
  var test register_data 
  json.NewDecoder(r.Body).Decode(&test) 
  fmt.Println(test)
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/auth/register", routes.Regsiter).Methods("POST")
  router.HandleFunc("/auth/login", routes.Login).Methods("POST")

  http.ListenAndServe(":8000", router)
}
