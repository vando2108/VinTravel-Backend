package auth

import (
  "net/http"
  "github.com/asaskevich/govalidator"
  "VinTravel/utils"
  "VinTravel/models"
) 

func Regsiter(writer http.ResponseWriter, request *http.Request) {
  username := request.PostFormValue("username")
  user_email := request.PostFormValue("user_email")
  password := request.PostFormValue("password")

  if govalidator.IsNull(username) || govalidator.IsNull(user_email) || govalidator.IsNull(password) {
    utils.JSON(writer, http.StatusBadRequest, "Data can not empty")
  }

  if !govalidator.IsEmail(user_email) {
    utils.JSON(writer, http.StatusBadRequest, "Email is invalid")
  }

  username = models.Sanitize(username)
  user_email = models.Sanitize(user_email)
  password = models.Sanitize(password)


}
