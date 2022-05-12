package main

import "database/sql"

func main(){
	connStr:="user=postgres password=1 dbname=fruitdb sslmode=disable"
	db,err:=sql.Open("postgres",connStr)
	
}