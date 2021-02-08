package repoimpl

import (
	models "VinTravel/models"
	repo "VinTravel/repository"
	"database/sql"
	"fmt"
)

type UserRepoImpl struct {
  Db *sql.DB
}

func NewUserRepo(db *sql.DB) repo.UserRepo {
  return &UserRepoImpl{
    Db: db,
  }
}

func (u *UserRepoImpl) Insert(user models.User) (error) {
  insertStament := `
    INSERT INTO users (id, uuid, name, user_email, password, created_at)
    VALUES ($1, $2, $3, $4, $5, $6)
  `

  _, err := u.Db.Exec(insertStament, 
  user.ID, user.Uuid, user.Name, user.User_email, user.Password, user.CreatedAt)

  if err != nil {
    panic(err)
  }

  fmt.Println("Add Succesfull: ", user)
  
  return nil
}

func (u *UserRepoImpl) UserExists(db *sql.DB, username string) (bool, error) {
  err := db.QueryRow("SELECT * FROM users WHERE username=?", username).Scan(&username)
  if err != nil {
    if err != sql.ErrNoRows {
      return false, err
    }
    return false, nil
  }
  return true, nil
}
