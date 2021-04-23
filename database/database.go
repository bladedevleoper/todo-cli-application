package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//construct the database credentials
func DbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_cli")
	if err != nil {
		panic(err.Error())
	}

	return db
}
