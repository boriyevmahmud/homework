package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Dbmethods struct {
}

func main() {
	connStr := "user=postgres password=1 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic("error while connecting to db")
	}
	defer db.Close()
	fmt.Println("1-Insert into Database\n2-Update Database(with ID)\n3-Delete Database(with ID)\n4-Get All")
	s := 0

	fmt.Scan(&s)
	switch s {
	case 1:
		{
			var (
				id        int64
				name      string
				phone     string
				email     string
				postalzip string
				country   string
			)
			fmt.Println("ID")
			fmt.Scan(&id)
			fmt.Println("Name")
			fmt.Scan(&name)
			fmt.Println("Phone")
			fmt.Scan(&phone)
			fmt.Println("Email")
			fmt.Scan(&email)
			fmt.Println("Postalzip")
			fmt.Scan(&postalzip)
			fmt.Println("Country")
			fmt.Scan(&country)

			dbs := Dbmethods{}

			usr, err := dbs.Create(Ecoreq{
				id:        id,
				name:      name,
				phone:     phone,
				email:     email,
				postalzip: postalzip,
				country:   country,
			})
			if err != nil {
				fmt.Println("error", err)
			}
			fmt.Println(usr, "added succesfully database")
			defer db.Close()
		}
	case 2:
		{
			var a, b string
			fmt.Println("Which ID you want update name of column ?")
			fmt.Scan(&a)
			fmt.Println("Name:")
			fmt.Scan(&b)
			fg := Dbmethods{}
			res := fg.Update(a, b)
			if res != nil {
				fmt.Println("error:", res)
			} else {
				fmt.Println("Successuflly updated Datas")
			}
			defer db.Close()
		}
	case 3:
		{
			var a string
			fmt.Println("Which ID you want remove information from Database:")
			fmt.Scan(&a)
			fg := Dbmethods{}
			res := fg.Delede(a)
			if res != nil {
				fmt.Println("error:", res)
			} else {
				fmt.Println("Successfully deleted information from Database")
			}
			defer db.Close()
		}
	case 4:
		{
			fg := Dbmethods{}
			res := fg.GetInfo()
			if res != nil {
				fmt.Println("error: ", res)
			}

		}
	}

}

func (h Dbmethods) Create(req Ecoreq) (Ecoreq, error) {
	user := &Ecoreq{}
	connStr := "user=postgres password=1 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("error")
		return Ecoreq{}, err
	}
	defer db.Close()

	query := `INSERT INTO	mytable (name,phone,email,postalzip,country) values ($1,$2,$3,$4,$5) returning *`
	err = db.QueryRow(query, req.name, req.phone, req.email, req.postalzip, req.country).Scan(
		&user.id,
		&user.name,
		&user.phone,
		&user.email,
		&user.postalzip,
		&user.country,
	)
	if err != nil {
		fmt.Println("error")
		return Ecoreq{}, err
	}
	return *user, nil

}
func (h Dbmethods) Update(a, b string) error {
	connStr := "user=postgres password=1 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("error")
		return err
	}
	defer db.Close()
	_, err = db.Exec(`UPDATE mytable SET name=$1 WHERE id=$2`, b, a)
	return err

}

func (h Dbmethods) Delede(a string) error {
	connStr := "user=postgres password=1 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("error")
		return err
	}
	defer db.Close()
	_, err = db.Exec(`DELETE FROM mytable WHERE id=$1`, a)
	return err

}

func (h Dbmethods) GetInfo() error {
	p := 0
	products := []Ecoreq{}
	fmt.Println("1-All Information\n2-Order by Name\n3-Datas where country=Turkey\n4-Datas where email @google.com")
	fmt.Scan(&p)
	connStr := "user=postgres password=1 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("error")
		return err
	}
	defer db.Close()
	rows, err := db.Query("select *from mytable")
	switch p {
	case 1:
		{
			rows, err = db.Query("select *from mytable")
			if err != nil {
				panic(err)
			}
		}
	case 2:
		{
			rows, err = db.Query("select *from mytable ORDER BY name")
			if err != nil {
				panic(err)
			}
		}
	case 3:
		{
			rows, err = db.Query(`SELECT *FROM mytable WHERE country='Turkey' `)
			if err != nil {
				panic(err)
			}
		}
	case 4:
		{
			rows, err = db.Query(`SELECT *FROM mytable WHERE email LIKE '%@google.com'`)
			if err != nil {
				panic(err)
			}
		}
	}

	for rows.Next() {
		s := Ecoreq{}
		err := rows.Scan(&s.id, &s.name, &s.phone, &s.email, &s.postalzip, &s.country)
		if err != nil {
			panic(err)
		}
		products = append(products, s)
	}
	defer rows.Close()
	for _, elem := range products {
		fmt.Println(elem.id, elem.name, elem.phone, elem.email, elem.postalzip, elem.country)
	}
	return err
}

type Methods interface {
	Create(Ecoreq, db string) (Ecoreq, error)
	Update(a, b string) error
	Delete(a string) error
	GetInfo(s int) (Ecoreq, error)
}

type Ecoreq struct {
	id        int64
	name      string
	phone     string
	email     string
	postalzip string
	country   string
}
