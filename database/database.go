package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

const TABLE = "todo_table"

func dbConnect(query string) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_cli")

	if err != nil {
		panic(err.Error())
	}

	insert, err := db.Query(query)

	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	//find out why defer is used
	defer insert.Close()
}

func InsertIntoTodo(data string) {
	query := "INSERT INTO " + TABLE + "(task_title, date_created) VALUES ('" + data + "', NOW())"
	dbConnect(query)
}

func GetTop10Tasks() {
	type Task struct {
		task_title   string
		created_date string
		id           int
	}
	statement := "SELECT id, task_title, date_created FROM " + TABLE + " WHERE task_title != '' ORDER BY id DESC LIMIT 10"
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_cli")

	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query(statement)

	defer rows.Close()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Here are the top 10 tasks")

	//itterate over each row
	for rows.Next() {
		//initialise a new task object
		task := Task{}

		//scan allocates the row column to its destination variable
		if err := rows.Scan(&task.id, &task.task_title, &task.created_date); err != nil {
			panic(err.Error())
		}

		fmt.Println(strconv.Itoa(task.id) + " - " + task.task_title + " - " + task.created_date)

	}
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
