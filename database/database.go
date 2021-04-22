package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const TABLE = "todo_table"

//construct the database credentials

func DbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_cli")
	if err != nil {
		panic(err.Error())
	}

	return db
}

func AddReminder(id int, date string) {
	type Reminder struct {
		todo_id          int
		date_of_reminder string
	}

	reminder := Reminder{}
	reminder.todo_id = id
	reminder.date_of_reminder = date

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_cli")

	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare(`INSERT INTO reminder_table (todo_id, date_of_reminder) VALUES (?,?);`)

	if err != nil {
		panic(err.Error())
	}

	_, es := stmt.Exec(reminder.todo_id, reminder.date_of_reminder)

	if es != nil {
		panic(es.Error())
	}

	fmt.Println("reminder added")

}
