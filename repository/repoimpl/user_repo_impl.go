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
    INSERT INTO users (id, uuid, name, username, password, created_at)
    VALUES ($1, $2, $3, $4, $5, $6)
  `

  _, err := u.Db.Exec(insertStament, 
  user.ID, user.Uuid, user.Name, user.Username, user.Password, user.CreatedAt)

  if err != nil {
    return err
  }

  fmt.Println("Add Succesfull: ", user)
  
  return nil
}

func (u *UserRepoImpl) NumberOfUsers() (int64, error) {
  sqlStmt := `
    SELECT COUNT(*) FROM users
  `
  rows, err := u.Db.Query(sqlStmt)
  if err != nil {
    return 0, err
  }
  var ret int64
  for rows.Next() {
    if err := rows.Scan(&ret); err != nil {
      return 0, nil
    }
  }

  return ret, err
}

func (u *UserRepoImpl) GetHashedPassword(username string) (string, error) {
  sqlStmt := " SELECT password FROM users WHERE username='" + username +"'"
  result, err := u.Db.Query(sqlStmt)
  if err != nil {
    return "", err
  }
  var password string
  for result.Next() {
    if err := result.Scan(&password); err != nil {
      return "", nil
    }
  }
  return password, err
}
