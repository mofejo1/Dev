package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL, 
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func main() {
	dbName := "users_database.db"

	// Fix 1: Open DB connection once and reuse it throughout main
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connection established")

	// Fix 2: Call createTable before inserting any users
	createTable(db)

	// Fix 3: Pass `db` (not dbName string) — createUser expects *sql.DB
	// Fix 4: Closed the unclosed string literal "john@doe.com"
	lastID, err := createUser(db, "John Doe", "john@doe.com", "password")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created user with ID:", lastID)

	// Fix 5: Fix invalid email "josephdoe.com" → "joseph@doe.com"
	lastID, err = createUser(db, "markDoe", "mark@doe.com", "password")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created user with ID:", lastID)

	lastID, err = createUser(db, "lukeDoe", "luke@doe.com", "password")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created user with ID:", lastID)

	lastID, err = createUser(db, "JosephDoe", "joseph@doe.com", "password")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created user with ID:", lastID)
}

// Fix 6: Return the error instead of calling log.Fatal — lets the caller decide
func createTable(db *sql.DB) error {
	_, err := db.Exec(schema)
	return err
}

// Fix 7: Rename parameter `hashedPassword` → `password` — it's plaintext coming in,
// hashing happens inside the function
func createUser(db *sql.DB, name, email, password string) (int64, error) {
	stmt := `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`

	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := db.Exec(stmt, name, email, string(hp))
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
