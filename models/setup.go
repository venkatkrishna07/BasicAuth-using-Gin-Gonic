package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

//This DB connection for other database func , not for authentication
func ConnectDatabase() {
	db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/TestAuth")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db

}
