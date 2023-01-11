package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(db *sql.DB) {
	fmt.Println("Database initialize")
	var sqlQuery string

	//Student table create
	sqlQuery = `CREATE TABLE if not exists user(
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"student_name" TEXT NOT NULL,
		"father_name" TEXT NOT NULL,
		"class" TEXT NOT NULL,
		"age" integer NOT NULL
	  );`

	log.Println("Student table created.")
	statement, err := db.Prepare(sqlQuery) // Prepare SQL Statement
	if err != nil {
		fmt.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
}
