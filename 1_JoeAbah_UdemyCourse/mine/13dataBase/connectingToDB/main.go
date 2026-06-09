package main

// This code is making first contact with a database. Like calling someone to check their phone is on before having a full conversation. No data is read or written yet — it just opens a connection, checks the database is alive, then closes it cleanly when done.
import (
	"database/sql" //instead this will be the one using the baove to acces the library
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // we are using a blank import becasuse we are not using anything from here directly
)

var schema = ` 
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL ,
	email TEXT NOT NULL UNIQUE,
	hashed_password BLOB NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func main() {
	dbName := "data.db"   // declare the name (just a string in memory, no file yet)
	_ = os.Remove(dbName) // delete any existing "data.db" file on disk (avoid collision)

	db, err := sql.Open("sqlite3", dbName) // now
	//  create a fresh "data.db" on disk
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("Closing database")
		if err := db.Close(); err != nil {
			log.Printf("error in closing data base: %v", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connection established")

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table was created")
}
