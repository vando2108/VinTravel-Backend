package main

import (
	// "github.com/julienschmidt/httprouter"
	// "net/http"
	"VinTravel/driver"
	"VinTravel/models"
	// repo "VinTravel/repository/repoimpl"
	"fmt"
	// "time"
)

const (
  host = "localhost"
  port = "5432"
  user = "postgres"
  password = "1234"
  dbname = "postgres"
)

func main() {
  db := driver.Connect(host, port, user, password, dbname)
  err := db.SQL.Ping()

  fmt.Println("Connect succesfull")

  if err != nil {
    panic(err)
  }

  pass, _ := models.Hash("Kd09112108")
  pass = models.Sanitize(pass)

  fmt.Println(models.CheckPasswordHash(pass, "Kd09112108"))
}
