package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Status string

const (
	Todo Status = "todo"
	InProgress Status = "in-progress"
	Done Status = "done"
)


type Task struct {
	Name string `json:"name"`
	ID int	`json:"id"`
	Description string	`json:"description"`
	Status Status	`json:"status"`
	CreatedAt time.Time	`json:"createdAt"`
	UpdatedAt time.Time	`json:"updatedAt"`
}


func commands() {
	arg := os.Args[1:]

	if len(arg) < 1 {
		help()
		return
	}

	switch arg[0] {
		case "add": if len(arg) < 2 {
			fmt.Println("Error: No task provided!")
			return
		}
		addTask(arg[1])

		case "update": if len(arg) < 2 {
			fmt.Println("Error: No full command provided")
			help()
			return
		}
		id, err := strconv.Atoi(arg[1])
		if err != nil {
			fmt.Println("Invalide task id provided",err)
			return
		}
		updateTask(id, arg[2])

		case "delete": if len(arg) < 2 {
			fmt.Println("Error: No task Provided!")
			return
			}
			id, err := strconv.Atoi(arg[1])
			if err != nil {
				fmt.Println("Invalide task id provided",err)
				return
			}
			deleteTask(id)
	}

}



func help() {
	fmt.Println("Usage: task-cli [command] [task]")
	fmt.Println("commands: ")
	fmt.Println("add - to add task")
}

func addTask(a string) {
	filePath := "task.json"

	var tasks []Task

	data, err := os.ReadFile(filePath)
	if err == nil {
		json.Unmarshal(data, &tasks)
	}

	newTask := Task{
		Name: a,
		ID: len(tasks) + 1,
		Status: Todo,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)
	

	jsontasks, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filePath, jsontasks, 0644)
	if err != nil {
		fmt.Println("Error writing file", err)
		return
	}

	fmt.Println("Added task:", newTask.Name)


}

func updateTask(id int, newName string) {
	filePath := "task.json"

	var tasks []Task

	data, err := os.ReadFile(filePath)
	if err == nil {
		json.Unmarshal(data, &tasks)
	}

	updated := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Name = newName
			tasks[i].UpdatedAt = time.Now()
			updated = true
			break
		}
	}

	if !updated {
		fmt.Println("Task not found with ID:", id)
		return
	}

		jsontasks, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	err = os.WriteFile(filePath, jsontasks, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Updated task", id, "to new name:", newName)
}

func deleteTask(id int) {
	var tasks []Task

	filePath := "task.json"

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
		return
	}


	deleted := false
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			deleted = true
			break
		}
	}

	if !deleted {
		fmt.Println("Index not found:", id)
	}

	for i := range tasks {
		tasks[i].ID = i + 1
	}

	jsontasks, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
	fmt.Println("Error marshaling JSON:", err)
	return
	}

	err = os.WriteFile(filePath, jsontasks, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Task deleted", id)
}


