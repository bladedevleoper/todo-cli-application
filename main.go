package main

import (
	//"fmt"
	"fmt"
	"github.com/bladedevleoper/go-cli-app/database"
	//"github.com/bladedevleoper/go-cli-app/handler"
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	ADDITEM      = 1
	ADDREMINDER  = 2
	TASKCOMPLETE = 3
)

var (
	firstEntry = false
	clear      map[string]func() //create a map for storing clear functions
)

const EMPTY = "Nothing has been enetered"

//will initialise the dynamically mapped clear functions
func init() {
	clear = make(map[string]func()) //initialise it
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

//TODO: Set up database schema for todo app (done)
//TODO: Set variables for user to input (done)
//TODO: Insert what the user has entered into the database with insert statement (done)
//TODO: Login before we access the todo app? (maybe??)
//TODO: display all tasks
//TODO: allow user to set a reminder on a task
//TODO: set task to completed
func main() {

	if firstEntry {
		callClearTerminal()
	}

	fmt.Print("Welcome to the TODO CLI Application: \n")
	fmt.Println("---------------------------------------")
	fmt.Println("Please select from the following menu:")
	fmt.Println("----------------------------------------")
	displayMenu()

	//will read user input
	reader := bufio.NewReader(os.Stdin)

	//infinate loop until we get text entered
	for {
		selection, _ := reader.ReadString('\n')

		//replace carriage and new line with blank string (this is specific to windows)
		selection = strings.Replace(selection, "\r\n", "", -1)

		//will handle the selected
		handleChosenSelection(selection)
	}
}

func callClearTerminal() {
	callFunction, ok := clear[runtime.GOOS] //get runtime os
	if ok {
		callFunction()
	} else {
		panic("Your platform is unsupported")
	}
}

func exitMessage() {
	fmt.Println("Thank you for using the TODO CLI App, chow!")
}

func callAddTask() {
	//should this be moved into a different package to handle these ?
	fmt.Print("What would you like to add? ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	for {
		if text != "" {
			database.InsertIntoTodo(text)
			fmt.Println("Item added")
			callAddTask()
		} else {
			fmt.Println(EMPTY)
			main()
		}
	}
}

func displayMenu() {
	menuList := []string{"Add Todo task", "Add Reminder", "Set task to complete"}
	listItem := 1
	for i := 0; i < len(menuList); i++ {
		fmt.Println(strconv.Itoa(listItem) + " - " + menuList[i])
		listItem++
	}

	fmt.Println("To Exit the Application - press ctrl + c")
	fmt.Println("Please Select a number from the menu")

}

func handleChosenSelection(selected string) {
	firstEntry = true
	switch selected {
	//if option 1 is selected from menu
	case strconv.Itoa(ADDITEM):
		callAddTask()
	case strconv.Itoa(ADDREMINDER):
		callAddReminder()
	default:
		exitMessage()
	}
}

//will set a reminder against a task
func callAddReminder() {
	//display top 10 recent tasks
	database.GetTop10Tasks()
	//ask a question, to the user on which task they would like to set a reminder against
	//then ask to set a date of the reminder
	//then insert into database
	//return back to main()
}
