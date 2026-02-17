package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// ------------------------------------
// structs
type Product struct{
	ID int `db:"id"`
	Name string `db:"name"`
	Price float64 `db:"price"`
	Stock int `db:"stock"`
}
type ProductStore struct{
	db *sqlx.DB
}
// ------------------------------------
// Connection to database
func ConnectToDatabase(dataSourceName string) (*ProductStore, error){
	db, err:= sqlx.Connect("postgres", dataSourceName)

	if err!=nil{
		return nil, err
	}

	return &ProductStore{db: db}, nil
}
// ------------------------------------




func main(){

	var connectQuery = "user=myuser password=mypassword dbname=mydb sslmode=disable"

	// connect to db
	var ps, err = ConnectToDatabase(connectQuery)
	fmt.Println(ps)
	fmt.Println(err)

}