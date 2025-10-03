package cmd

import (
	"fmt"
	"github.com/Krish-jswl/task-cli-go/internal/storage"
)

func DeleteTask(id int) {

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
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
		return
	}

	for i := range tasks {
		tasks[i].ID = i + 1
	}

	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Task deleted", id)
}
