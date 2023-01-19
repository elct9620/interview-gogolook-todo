package todo

import "github.com/elct9620/interview-gogolook-todo/internal/todo/domain"

type TodoUseCase struct {
}

func NewTodoUseCase() *TodoUseCase {
	return &TodoUseCase{}
}

func (uc *TodoUseCase) ListTasks() []map[string]any {
	task := domain.NewTask(1, "name", false)

	return []map[string]any{
		{"id": task.ID, "name": task.Name, "status": task.StatusCode()},
	}
}
