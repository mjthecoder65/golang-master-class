package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// The underscore import is used to import a package solely for its side-effect,
	// Specifically for its initialization.
	// This technique is often used when you need to register a package's functionality
	// without directly referencing it in your code.
	_ "github.com/lib/pq"
)

/*
 Learning about how to connect and interacting with Databases.
 I will use Postgreql.
  Requirements:
	1. pq driver: A driver for go to interact with PostgreSQL
	2. go get github.com/lib/pq
	3. DB connection string.
*/

type DBConfig struct {
	Username string
	Password string
	DbName   string
	Hostname string
	SSLMode  string
	Port     int
}

func CreateTodoTable(db *sql.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS todos (
		 id SERIAL PRIMARY KEY,
		 name VARCHAR(100) NOT NULL,
		 is_complete BOOLEAN DEFAULT  FALSE
		)
	`
	_, err := db.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table SQL created successfully!")
}

func AddTodo(name string, db *sql.DB) {
	insertSQL := `INSERT INTO todos (name) VALUES ($1)`

	_, err := db.Exec(insertSQL, name)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Todo added successfully!")
}

func GetTodos(db *sql.DB) []string {
	query := "SELECT id, name, is_complete FROM todos"

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	todos := make([]string, 0, 10)

	for rows.Next() {
		var id int
		var name string
		var is_completed bool

		err = rows.Scan(&id, &name, &is_completed)

		for err != nil {
			log.Fatal(err)
		}

		todos = append(todos, name)
	}

	return todos
}

// init function is automatically called when this package is imported.
func init() {
	fmt.Println("Package initialization")
}

func LearnDBConnection() {
	config := DBConfig{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DB"),
		Hostname: os.Getenv("POSTGRES_HOST"),
		SSLMode:  "disable",
		Port:     5030,
	}

	connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s", config.Username,
		config.Password, config.DbName, config.Hostname, config.Port, config.SSLMode)

	fmt.Println(connectionStr)

	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Test the connection
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database....")

	// Create a table
	CreateTodoTable(db)

	// todos := []string{"Learn About goroutines", "Clean my room", "Do 20min Exercise"}

	// for _, todo := range todos {
	// 	AddTodo(todo, db)
	// }

	todos := GetTodos(db)

	for _, todo := range todos {
		fmt.Println(todo)
	}
}
