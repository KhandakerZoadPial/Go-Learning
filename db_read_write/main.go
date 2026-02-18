package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// ------------------------------------
// structs
type Product struct {
	ID    int     `db:"id"`
	Name  string  `db:"name"`
	Price float64 `db:"price"`
	Stock int     `db:"stock"`
}
type ProductStore struct {
	db *sqlx.DB
}

// ------------------------------------
// Connection to database
func ConnectToDatabase(dataSourceName string) (*ProductStore, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}

	return &ProductStore{db: db}, nil
}

// ------------------------------------
// add data to database
func (ps *ProductStore) AddProduct(name string, price float64, stock int) error {
	query := `INSERT INTO products (name, price, stock) VALUES ($1, $2, $3)`
	_, err := ps.db.Exec(query, name, price, stock)
	return err
}

// fetch all data
func (ps *ProductStore) GetAllProduct() ([]Product, error) {
	var products []Product
	query := `SELECT id,name,price, stock FROM products`
	err := ps.db.Select(&products, query)

	return products, err
}

// fetch single data
func (ps *ProductStore) GetAProduct(id int) (Product, error) {
	var p Product
	query := `SELECT id,name,price,stock FROM products WHERE id=$1`
	err := ps.db.Get(&p, query, id)
	return p, err
}

// delete a record
func (ps *ProductStore) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id=$1`
	_, err := ps.db.Exec(query, id)
	return err
}

func main() {

	var connectQuery = "user=myuser password=mypassword dbname=mydb sslmode=disable"

	// connect to db
	var ps, err = ConnectToDatabase(connectQuery)
	fmt.Println(ps)
	fmt.Println(err)

	if err != nil {
		fmt.Println("Connection error:", err)
	}

	ps.db.Close()

	// fmt.Println("Adding Products.....")
	// ps.AddProduct("Keyboard", 150.0, 5)
	// ps.AddProduct("Mouse", 80.0, 15)

	// allProducts, err := ps.GetAllProduct()
	fmt.Println("Before deleting id 2")
	singleProduct, err := ps.GetAProduct(2)
	fmt.Println(singleProduct, err)
	ps.DeleteProduct(2)
	fmt.Println("After deleting id 2")
	singleProduct, err = ps.GetAProduct(2)
	fmt.Println(singleProduct, err)

}
