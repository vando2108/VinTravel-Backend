package repository

import (
	"VinTravel/models"
	"database/sql"
)

type UserRepo interface {
  Insert(user models.User) (error)
  UserExists(db *sql.DB, username string) (bool, error)
}
