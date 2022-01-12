package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Task struct {
	Id          int
	Task     string
	IsFinished bool
}

var TaskList = []Task{}

func main() {
	TaskList = make([]Task, 0)
	serverRun()
}

// Starting the server on port 8080 and catching "/"" and "/add"
func serverRun() {
	fmt.Println("Running server...")

	http.HandleFunc("/add", addTask)
	http.HandleFunc("/", taskList)
	http.ListenAndServe(":8080", nil)
}

// Display function of all tasks
func taskList(w http.ResponseWriter, r *http.Request) {
	taskValue := r.PostFormValue("Task")
	if taskValue != "" {
		fmt.Println("Task:", taskValue)
		t := &Task{Id: len(TaskList) + 1, Task: taskValue, IsFinished: false}
		TaskList = append(TaskList, *t)
	}

	t, err := template.ParseFiles("templates/main.gtpl")
	if err != nil {
		log.Fatal("Can not parse templates/main.gtpl " + err.Error())
	}
	t.Execute(w, TaskList)
}

// Function for adding a new task
func addTask(w http.ResponseWriter, r *http.Request) {
	log.Print("addTask")
	t, _ := template.ParseFiles("templates/add.gtpl")
	t.Execute(w, t)
}
