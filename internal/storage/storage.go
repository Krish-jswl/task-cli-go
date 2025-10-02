package storage

import(
	"os"
	"encoding/json"
	"github.com/Krish-jswl/task-cli-go/dtask"
)
const taskFile = "task.json"

func LoadTasks() ([]dtask.Task, error) {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []dtask.Task{}, nil
		}
		return nil, err
	}

	var tasks []dtask.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(tasks []dtask.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(taskFile, data, 0644)
}
