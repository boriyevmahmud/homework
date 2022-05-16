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
	

	query := `select * from users`
	result := db.QueryRow(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	
}
