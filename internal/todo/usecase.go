package todo

import "github.com/elct9620/interview-gogolook-todo/internal/todo/domain"

type TaskOutput struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type TodoRepository interface {
	AllTasks() []*domain.Task
}

type TodoUseCase struct {
	repo TodoRepository
}

func NewTodoUseCase(repo TodoRepository) *TodoUseCase {
	return &TodoUseCase{repo}
}

func (uc *TodoUseCase) ListTasks() (output []TaskOutput) {
	output = make([]TaskOutput, 0)

	tasks := uc.repo.AllTasks()
	for idx := range tasks {
		output = append(output, TaskOutput{
			ID:     tasks[idx].ID,
			Name:   tasks[idx].Name,
			Status: tasks[idx].StatusCode(),
		})
	}

	return output
}

func (uc *TodoUseCase) CreateTask(name string) TaskOutput {
	return TaskOutput{1, name, 0}
}
