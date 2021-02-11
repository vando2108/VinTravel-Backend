package models

import (
	"html"
	"strings"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
  ID int64 `json:"id"`
  Uuid string `json:"uuid"`
  Name string `json:"name"` //is display name
  Username string `json:"username"`
  Password string `json:"password"`
  CreatedAt time.Time `json:"created_at"`
}

func Hash(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
  return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Sanitize(data string) string {
  data = html.EscapeString(strings.TrimSpace(data))
  return data
}
