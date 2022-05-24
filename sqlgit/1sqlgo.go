package main

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

func main() {
  connStr := "user=postgres password=1 dbname=testdb sslmode=disable"
  db, err := sql.Open("postgres", connStr)

  if err != nil {
    panic("error while connecting to db")
  }
  defer db.Close()

  query := `INSERT INTO users (username, pasword, email) values($1,$2,$3) returning id`
  result, err := db.Exec(query, "Rustam", "aaaaa", "rustik@gmail.com")
  if err != nil {
    fmt.Println(err)
    return
  }
  

  fmt.Println(result.LastInsertId())
  fmt.Println(result.RowsAffected())
}