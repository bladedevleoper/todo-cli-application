package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const TABLE = "todo_table"

func DbConnect(data string) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_cli")

	if err != nil {
		panic(err.Error())
	}

	insert, err := db.Query(insertIntoTodo(data))

	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	//find out why defer is used
	defer insert.Close()
}

func insertIntoTodo(data string) string {
	return "INSERT INTO " + TABLE + "(task_title, date_created) VALUES ('" + data + "', NOW())"
}
