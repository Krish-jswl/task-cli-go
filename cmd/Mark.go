package cmd

import (
	"fmt"
	"github.com/Krish-jswl/task-cli-go/internal/storage"
	"github.com/Krish-jswl/task-cli-go/dtask"
)

func Mark(id int, status string) {


	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	
	for i, task := range tasks {
		if task.ID == id {
			// tasks[i].Status = Done
			switch status{
			case "mark-done":
				tasks[i].Status = dtask.Done
				fmt.Printf("task %d marked done\n", id)
			case "mark-todo":
				tasks[i].Status = dtask.Todo
				fmt.Printf("task %d marked todo\n", id)
			case "mark-in-progress":
				tasks[i].Status = dtask.InProgress
				fmt.Printf("task %d marked in-progress\n", id)
			default:
				fmt.Println("Invalide command provided")
				Help()
			}
			break
		}
	}

	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

}

