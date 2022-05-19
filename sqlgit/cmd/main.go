package main

import (
	"database/sql"
	"fmt"

	//"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// 1. database/sql - CRUD
// 2. GORM
// 3. array postgres example: postal_codes, phone_number

type Customer struct {
	ID          string
	FirstName   string
	LastName    string
	Username    string
	//PhoneNumber []Phone
	//Adress      []Adress
	//Products    []Product
	Email       string
	Gender      string
	Birthdate   string
	Password    string // should be hashed and validate password should be 8 symbols
	Status      string
}

type Adress struct {
	ID          string
	Customer_id string
	Country     string
	City        string
	District    string
	PostalCodes int64
}

type Product struct {
	ID          string
	Customer_id string
	Name        string
	Types       Type
	Cost        int64
	OrderNumber int64
	Amount      int64
	Currency    string
	Rating      int64
}

type Phone struct {
	ID          string
	Customer_id string
	Numbers     int64
	Code        string
}

type Type struct {
	ID         int64
	Product_id string
	Name       string
}

type Dbmethods struct {
}

func main() {
	connStr := "user=postgres password=1 dbname=shop sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	check(err)
	defer db.Close()
	s := 0
	fmt.Println("1-Insert Into\n2-Update UUID\n3-Delete With UUID\n4-Get ALL")
	fmt.Scan(&s)
	switch s {
	case 1:
		{

			var (
				FirstName string
				LastName  string
				Username  string
				Email     string
				Gender    string
				Birthdate string
				Password  string
				Status    string

				Country     string
				City        string
				District    string
				PostalCodes int64

				Numbers int64  // 998812891, 998802891 ...
				Code    string // "+998"

				Name_product string
				Cost         int64
				OrderNumber  int64
				Amount       int64
				Currency     string
				Rating       int64

				Name_type string
			)
			fmt.Println("FirstName")
			fmt.Scan(&FirstName)
			fmt.Println("LastName")
			fmt.Scan(&LastName)
			fmt.Println("Username")
			fmt.Scan(&Username)
			fmt.Println("Adress\nCountry")
			fmt.Scan(&Country)
			fmt.Println("City")
			fmt.Scan(&City)
			fmt.Println("District")
			fmt.Scan(&District)
			fmt.Println("Postalcodes")
			fmt.Scan(&PostalCodes)
			fmt.Println("Phone\nNumbers")
			fmt.Scan(&Numbers)
			fmt.Println("Code")
			fmt.Scan(&Code)
			fmt.Println("Product\nName_product")
			fmt.Scan(&Name_product)
			fmt.Println("Product type")
			fmt.Scan(&Name_type)
			fmt.Println("Cost")
			fmt.Scan(&Cost)
			fmt.Println("Order_Number")
			fmt.Scan(&OrderNumber)
			fmt.Println("Amount")
			fmt.Scan(&Amount)
			fmt.Println("Currency")
			fmt.Scan(&Currency)
			fmt.Println("Rating")
			fmt.Scan(&Rating)
			fmt.Println("Email")
			fmt.Scan(&Email)
			fmt.Println("Gender")
			fmt.Scan(&Gender)
			fmt.Println("Birthday")
			fmt.Scan(&Birthdate)
			fmt.Println("Password")
			fmt.Scan(&Password)
			fmt.Println("Status")
			fmt.Scan(&Status)

			dbs := Dbmethods{}
			err = dbs.Create(Customer{
				FirstName: FirstName,
				LastName:  LastName,
				Username:  Username,
				Email:     Email,
				Gender:    Gender,
				Birthdate: Birthdate,
				Password:  Password,
				Status:    Status},
				Adress{
					Country:     Country,
					City:        City,
					District:    District,
					PostalCodes: PostalCodes},
				Product{
					Name:        Name_product,
					Cost:        Cost,
					OrderNumber: OrderNumber,
					Amount:      Amount,
					Currency:    Currency,
					Rating:      Rating},
				Phone{
					Numbers: Numbers,
					Code:    Code},
				Type{
					Name: Name_type,
				},
			)
			if err == nil {
				fmt.Println("Informations added successfully to DATABASE")
			}
		}
	case 2:
		{

		}
	case 3:
		{
			var a string
			fmt.Println("Enter the UUID of customer:")
			fmt.Scan(&a)
			dbs := Dbmethods{}
			err := dbs.Delede(a)
			check(err)
			fmt.Println("Information deleted from Database successfully")
		}
	case 4:
		{
			dbs := Dbmethods{}
			err := dbs.Getall()
			check(err)

		}
	}

}

func (h Dbmethods) Create(cus Customer, ad Adress, pr Product, ph Phone, ty Type) error {
	var idd, idd1 string
	connStr := "user=postgres password=1 dbname=shop sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("1poing", err)
	}
	defer db.Close()

	query := `INSERT INTO	customers (firstname,lastname,username,email,gender,birthdate,password,status) values ($1,$2,$3,$4,$5,$6,$7,$8) returning customer_id`
	err = db.QueryRow(query, cus.FirstName, cus.LastName, cus.Username, cus.Email, cus.Gender, cus.Birthdate, cus.Password, cus.Status).Scan(&idd)
	//check(err)
	if err != nil {
		fmt.Println("2-point", err)
	}

	query = `insert into adress(customer_id,country,city,district,postalcodes) 
	values($1,$2,$3,$4,$5)`
	_, err = db.Exec(query, idd, ad.Country, ad.City, ad.District, ad.PostalCodes)
	//check(err)
	if err != nil {
		fmt.Println("3-point", err)
	}
	query = `insert into product(customer_id,name,cost,ordernumber,amount,currency,rating)
	values($1,$2,$3,$4,$5,$6,$7) returning product_id`
	err = db.QueryRow(query, idd, pr.Name, pr.Cost, pr.OrderNumber, pr.Amount, pr.Currency, pr.Rating).Scan(&idd1)
	//check(err)
	if err != nil {
		fmt.Println("4-point", err)
	}
	query = `insert into phone(customer_id,numbers,code)
	values($1,$2,$3)`
	_, err = db.Exec(query, idd, ph.Numbers, ph.Code)
	//check(err)
	if err != nil {
		fmt.Println("5-point", err)
	}
	query = `insert into typee(product_id,name)values($1,$2)`
	_, err = db.Exec(query, idd1, ty.Name)
	check(err)

	return err
}

func (h Dbmethods) Delede(a string) error {
	connStr := "user=postgres password=1 dbname=shop sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("error")
		return err
	}
	defer db.Close()
	_, err = db.Exec(`DELETE FROM customers WHERE id=$1`, a)
	defer db.Close()
	return err
}

func (h Dbmethods) Getall() error {
	cus := []Customer{}
	ph := []Phone{}
	add := []Adress{}
	pro := []Product{}
	//ptt := []Type{}

	connStr := "user=postgres password=1 dbname=shop sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	check(err)
	defer db.Close()
	rows, err := db.Query(`select c.customer_id,firstname,lastname,username,country,city,district,postalcodes,
			numbers,code,p.name,cost,ordernumber,amount,currency,rating,email,gender,birthdate,
			password,status from customers c inner join adress a on c.customer_id=a.customer_id inner join 
			product p on a.customer_id=p.customer_id inner join phone h on p.customer_id=h.customer_id inner join 
			typee t on p.product_id=t.product_id  ;`)
	check(err)
	for rows.Next() {
		c := Customer{}
		p := Phone{}
		ad := Adress{}
		pr := Product{}
		//pt:=Type{}
		err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Username, &ad.Country, &ad.City, &ad.District, &ad.PostalCodes, &p.Numbers, &p.Code, &pr.Name, &pr.Cost, &pr.OrderNumber, &pr.Amount, &pr.Currency, &pr.Rating, &c.Email, &c.Gender, &c.Birthdate, &c.Password, &c.Status)
		check(err)
		cus = append(cus, c)
		add = append(add, ad)
		ph = append(ph, p)
		pro = append(pro, pr)
	}
	defer db.Close()
	for i := 0; i < len(cus); i++ {
		fmt.Print(cus[i], ph[i], pro[i])
		for j := i; j < i+1; j++ {
			fmt.Print(add[j])
		}
		for k := i; k < i+1; k++ {
			fmt.Print(ph[k])
		}
		for l := i; l < i+1; l++ {
			fmt.Print(pro[l])
		}
		fmt.Print("\n")
	}
	return err
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
