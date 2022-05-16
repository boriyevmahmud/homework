package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

func main() {
	connStr := "user=postgres password=1 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic("error while connecting to db")
	}
	defer db.Close()
	// var id int
	// user := User{
	// 	Username: "Ramazon",
	// 	Password: "20031031",
	// 	Email:    "ramazonergashev31@gmail.com",
	// }

	// query := `INSERT INTO users (username, pasword, email) values($1,$2,$3) returning id`
	// err = db.QueryRow(query, user.Username, user.Password, user.Email).Scan(&id)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	urs := []User{}
	execQuery := `UPDATE users SET email=$1 WHERE id=$2`
	_, err = db.Exec(execQuery, "newemail2@gmail.com", 2)

	if err != nil {
		fmt.Println(err)
		return
	}
	queryUpdate := `SELECT id,username,pasword,email FROM users`
	rows, err := db.Query(queryUpdate)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		us := User{}
		err := rows.Scan(&us.ID, &us.Username, &us.Password, &us.Email)
		if err != nil {
			fmt.Println(err)
			return
		}
		urs = append(urs, us)
	}
	for _, u := range urs {
		fmt.Println(u)
	}

}
