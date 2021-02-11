package routes

import (
	"VinTravel/configs"
	"VinTravel/driver"
	"VinTravel/models"
	repo "VinTravel/repository/repoimpl"
	"VinTravel/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/asaskevich/govalidator"
) 

type register_data struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Name string `json:"name"`
}

func Regsiter(writer http.ResponseWriter, request *http.Request) {
  decoder := json.NewDecoder(request.Body)
  var FormData register_data
  _ = decoder.Decode(&FormData)
  username := FormData.Username
  password := FormData.Password
  name := FormData.Name

  if govalidator.IsNull(username)|| govalidator.IsNull(password) {
    utils.JSON(writer, http.StatusBadRequest, "Data can not empty")
    return
  }

  username = models.Sanitize(username)
  password = models.Sanitize(password)

  db := driver.Connect(configs.DB_HOST, configs.DB_PORT, configs.DB_USER, configs.DB_PASSWORD, configs.DB_NAME)
  pingTest := db.SQL.Ping()
  if pingTest != nil {
    fmt.Println("Regsiter error: cannot access database")
    utils.JSON(writer, http.StatusInternalServerError, "Cannot access database")
    return
  }
  fmt.Println("Connect Succesfully")
  
  var err error
  password, err = models.Hash(password)
  userRepo := repo.NewUserRepo(db.SQL)
  
  if err != nil {
    fmt.Println("Hash error")
    utils.JSON(writer, http.StatusInternalServerError, "Can not hash password")
    return
  }

  newId, _ := userRepo.NumberOfUsers()
  user := models.User {
    ID: newId,
    Uuid: strconv.Itoa(rand.Intn(10000) + 1000),
    Name: name,
    Username: username,
    Password: password,
    CreatedAt: time.Now(),
  }

  err = userRepo.Insert(user)
  if err != nil {
    fmt.Println("Insert new user error: ", err)
    utils.JSON(writer, http.StatusInternalServerError, "User does exists")
    return
  }

  utils.JSON(writer, http.StatusCreated, "Register Succesfully")
  fmt.Println("Regsiter Succesfully")
}

type login_data struct {
  Username string `json:"username"`
  Password string `json:"password"`
}

func Login(writer http.ResponseWriter, request *http.Request) {
  var FormData login_data
  json.NewDecoder(request.Body).Decode(&FormData)
  if govalidator.IsNull(FormData.Username) || govalidator.IsNull(FormData.Password) {
    utils.JSON(writer, http.StatusBadRequest, "Data can not empty")
    return
  }

  FormData.Username = models.Sanitize(FormData.Username)
  FormData.Password = models.Sanitize(FormData.Password)

  db := driver.Connect(configs.DB_HOST, configs.DB_PORT, configs.DB_USER, configs.DB_PASSWORD, configs.DB_NAME)
  pingTest := db.SQL.Ping()
  if pingTest != nil {
    fmt.Println("Login error: cannot access database")
    utils.JSON(writer, http.StatusInternalServerError, "Cannot access database")
    return
  }
  fmt.Println("Connect Succesfully")

  userRepo := repo.NewUserRepo(db.SQL)
  
  hashedPassword, err := userRepo.GetHashedPassword(FormData.Username)
  if err != nil {
    utils.JSON(writer, http.StatusInternalServerError, "User dose not exists")
    fmt.Println("Login error: User dose not exists")
    return
  }

  if models.CheckPasswordHash(hashedPassword, FormData.Password) == nil {
    utils.JSON(writer, http.StatusAccepted, "Login accepted")
    fmt.Println("Login has accepted: ", FormData)
    return
  }
  utils.JSON(writer, http.StatusUnauthorized, "Login Unauthorized")
}
