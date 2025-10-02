package cmd

import(
	"fmt"
)


func Help() {
	fmt.Println("Usage: task-cli [command] [task]")
	fmt.Println("commands: ")
	fmt.Println("add - to add task")
}

