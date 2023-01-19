package todo

import "github.com/elct9620/interview-gogolook-todo/internal/todo/domain"

type TaskOutput map[string]any

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
			"id":     tasks[idx].ID,
			"name":   tasks[idx].Name,
			"status": tasks[idx].StatusCode(),
		})
	}

	return output
}
