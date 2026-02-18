package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Middleware-----------------------------------------------------------
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("LOG: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
	}
}

// structs--------------------------------------------------------------
type Todo struct {
	ID       int    `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	IsDone   bool   `db:"is_done" json:"is_done"`
	Priority int    `db:"priority" json:"priority"`
}

type TodoStore struct {
	db *sqlx.DB
}

// -------------------------------------------------------------------
// Connection to database
func ConnectToDatabase(dataSourceName string) (*TodoStore, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}

	return &TodoStore{db: db}, nil
}

// CRUDS----------------------------------------------------------------
func (ts *TodoStore) AddTodo(title string, priority int) (int, error) {
	var id int
	query := "INSERT INTO todos (title,is_done,priority) VALUES ($1, $2, $3)"
	err := ts.db.QueryRowx(query, title, false, priority).Scan(&id)
	return id, err

}

func (ts *TodoStore) GetTodos() ([]Todo, error) {
	var data []Todo
	query := `SELECT id, title, is_done, priority FROM todos`
	err := ts.db.Select(&data, query)

	return data, err
}

// func (ts *TodoStore) MarkAsDone(id int) error {

// }

// ENDPOINTS-------------------------------------------------------------
func (ts *TodoStore) createTodo(w http.ResponseWriter, r *http.Request) {
	var data Todo
	err := json.NewDecoder(r.Body).Decode(&data)

	ts.AddTodo(data.Title, data.Priority)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (ts *TodoStore) getAllTodo(w http.ResponseWriter, r *http.Request) {
	data, err := ts.GetTodos()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {

	var connectQuery = "user=myuser password=mypassword dbname=mydb sslmode=disable"

	// connect to db
	ts, err := ConnectToDatabase(connectQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("POST /create_todo", loggingMiddleware(ts.createTodo))
	http.HandleFunc("GET /get_all_todos", loggingMiddleware(ts.getAllTodo))
	fmt.Println("Server starting on: 8080...")
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
