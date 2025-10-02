package cmd

import (
	"github.com/Krish-jswl/task-cli-go/internal/storage"
	"fmt"
	"time"
)

func UpdateTask(id int, newName string) {

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
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

	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Updated task", id, "to new name:", newName)
}

