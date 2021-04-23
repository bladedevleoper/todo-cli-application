package models

import (
	"github.com/bladedevleoper/go-cli-app/database"
)

type Reminder struct {
	todo_id          int
	date_of_reminder string
}

func AddReminder(id int, date string) {

	reminder := Reminder{}
	reminder.todo_id = id
	reminder.date_of_reminder = date

	db := database.DbConnect()

	stmt, err := db.Prepare(`INSERT INTO reminder_table (todo_id, date_of_reminder) VALUES (?,?);`)

	if err != nil {
		panic(err.Error())
	}

	_, es := stmt.Exec(reminder.todo_id, reminder.date_of_reminder)

	if es != nil {
		panic(es.Error())
	}
}
