package driver

import (
  "fmt"
  _ "github.com/lib/pq"
  "database/sql"
)

type PostgresDB struct {
  SQL *sql.DB
} 

var Postgres = &PostgresDB{}

func Connect(host, port, user, password, dbname string) (*PostgresDB) {
  connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				host, port, user, password, dbname)
  db, err := sql.Open("postgres", connectionStr)
  if err != nil {
    panic(err)
  }

  Postgres.SQL = db
  return Postgres
}
