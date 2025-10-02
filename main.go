package main

import (
	"os"
	"fmt"
	"github.com/Krish-jswl/task-cli-go/cmd"
	"strconv"
)

func main() {
	
	arg := os.Args[1:]

	if len(arg) < 1 {
		cmd.Help()
		return
	}

	switch arg[0] {
		case "add": if len(arg) < 2 {
			fmt.Println("Error: No task provided!")
			return
		}
		cmd.AddTask(arg[1])

		case "update": if len(arg) < 2 {
			fmt.Println("Error: No full command provided")
			cmd.Help()
			return
		}
		id, err := strconv.Atoi(arg[1])
		if err != nil {
			fmt.Println("Invalide task id provided",err)
			return
		}
		cmd.UpdateTask(id, arg[2])

		case "delete": if len(arg) < 2 {
			fmt.Println("Error: No task Provided!")
			return
			}
			id, err := strconv.Atoi(arg[1])
			if err != nil {
				fmt.Println("Invalide task id provided",err)
				return
			}
			cmd.DeleteTask(id)
		
		case "mark-done", "mark-todo", "mark-in-progress":	if len(arg) < 2 {
			fmt.Println("Error: No task Provided!")
			return
			}
			id, err := strconv.Atoi(arg[1])
			if err != nil {
				fmt.Println("Invalide task id provided",err)
				return
			}
			cmd.Mark(id, arg[0])

		case "list":
		var status string
		if len(arg) > 1 {
			status = arg[1] 
		}
	cmd.List(status)
	}

}
