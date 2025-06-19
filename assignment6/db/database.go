package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "local.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	statement, err := db.Prepare(`CREATE TABLE if not exists employees (
		employee_id integer,		
		employee_name varchar(255),
		employee_designation varchar(255),
		employee_salery REAL
	  )`)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()

	defer db.Close()

	return db, nil
}

// func main() {
// 	db, err := InitDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(db)
// }
