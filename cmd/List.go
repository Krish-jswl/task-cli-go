package cmd

import (
	"fmt"
	"github.com/Krish-jswl/task-cli-go/internal/storage"
	"github.com/Krish-jswl/task-cli-go/dtask"
) 

func List(status string) {

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	fmt.Println("")
	fmt.Printf("%-5s %-30s %-15s %-15s\n", "ID", "Task-Name", "Status", "Created Date")
	fmt.Println("------------------------------------------------------------------------")
	for _, task := range tasks {
			if status == "" || task.Status == dtask.Status(status) {
					fmt.Printf("%-5d %-30s %-15s %-15s\n", task.ID, task.Name, task.Status, task.CreatedAt.Format("02-Jan-2006"))
			}
	}
	fmt.Println("")
}

