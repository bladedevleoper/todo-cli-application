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
	statement := "SELECT id, task_title, date_created FROM " + TABLE + "  ORDER BY id DESC LIMIT 10"
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
