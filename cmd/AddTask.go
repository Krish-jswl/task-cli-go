package cmd

import (
	"fmt"
	"github.com/Krish-jswl/task-cli-go/internal/storage"
	"github.com/Krish-jswl/task-cli-go/dtask"
	"time"
)
func AddTask(a string) {

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	newTask := dtask.Task{
		Name: a,
		ID: len(tasks) + 1,
		Status: dtask.Todo,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)
	

	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Added task:", newTask.Name)

}

