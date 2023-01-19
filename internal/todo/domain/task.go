package domain

type TaskStatus int

const (
	TaskIncomplete TaskStatus = iota
	TaskComplete
)

type Task struct {
	ID     int
	Name   string
	Status TaskStatus
}

func NewTask(id int, name string, status TaskStatus) *Task {
	return &Task{
		ID:     id,
		Name:   name,
		Status: status,
	}
}
