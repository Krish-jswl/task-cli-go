# Task Manager Cli written in Golang
- for : https://roadmap.sh/projects/task-tracker

## Installation
```sh
git clone https://github.com/Krish-jswl/task-cli-go.git
cd task-cli-go/
go build .
```

## Commands
```bash
./task-cli-go add [task] # To add task with some name
./task-cli-go update [id] [new-task-name] # To update a specific task with new name
./task-cli-go list # Show all task
./task-cli-go list [status] # Show all task with specific status , [status] = todo/done/in-progress 
./task-cli-go mark-[status] [id] # To mark a task at specific index with new status , [status] = todo/done/in-progress
./task-cli-go delete [id] # To delete a task
```
