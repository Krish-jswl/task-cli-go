package cmd

import(
	"fmt"
)


func Help() {
	fmt.Println("")
	fmt.Println("Usage: ")
	fmt.Println("./task-cli-go add [\"task\"]")
	fmt.Println("./task-cli-go list")
	fmt.Println("./task-cli-go list [status]")
	fmt.Println("./task-cli-go mark-[status] [id]", )
	fmt.Println("./task-cli-go delete [id]")
	fmt.Println("")
	fmt.Println("[status] = todo/done/in-progress")
	fmt.Println("")
}

