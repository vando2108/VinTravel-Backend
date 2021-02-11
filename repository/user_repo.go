package repository

import (
	"VinTravel/models"
)

type UserRepo interface {
  Insert(user models.User) (error)
  NumberOfUsers() (int64, error)
  GetHashedPassword(username string) (string, error) 
}
