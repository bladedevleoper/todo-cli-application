package models

import (
	"fmt"
	"github.com/bladedevleoper/go-cli-app/database"
	"strconv"
)

const TABLE = "todo_table"

var (
	db = database.DbConnect()
)

func InsertTask(payload string) {
	query := "INSERT INTO " + TABLE + "(task_title, date_created) VALUES ('" + payload + "', NOW())"
	insert, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func GetTop10Tasks() {
	type Task struct {
		task_title   string
		created_date string
		id           int
	}
	statement := "SELECT id, task_title, date_created FROM " + TABLE + " WHERE task_title != '' ORDER BY id DESC LIMIT 10"

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
