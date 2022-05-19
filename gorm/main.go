package main
import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	
)

func main(){
	connStr:="user=postgres password=1 dbname=testdb sslmode=disable"
	db,err:=gorm.Open(sql.Open("postgres",connStr),&gorm.Config{})
	if err!=nil{
		fmt.Println(err,"Point-1")
	}
}