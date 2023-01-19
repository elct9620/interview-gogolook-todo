package todo

type TodoUseCase struct {
}

func NewTodoUseCase() *TodoUseCase {
	return &TodoUseCase{}
}

func (uc *TodoUseCase) ListTasks() []map[string]any {
	return []map[string]any{
		{"id": 1, "name": "name", "status": 0},
	}
}
