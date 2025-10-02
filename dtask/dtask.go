package dtask

import (
	"time"
)

type Status string

const (
	Todo Status = "todo"
	InProgress Status = "in-progress"
	Done Status = "done"
)


type Task struct {
	Name string `json:"name"`
	ID int	`json:"id"`
	Description string	`json:"description"`
	Status Status	`json:"status"`
	CreatedAt time.Time	`json:"createdAt"`
	UpdatedAt time.Time	`json:"updatedAt"`
}

