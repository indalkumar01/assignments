package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "local.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	statement, err := db.Prepare(`CREATE TABLE employees (
		id integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		employee_name varchar(255)		
	  )`)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()

	defer db.Close()

	return db, nil
}

func main() {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)
}
