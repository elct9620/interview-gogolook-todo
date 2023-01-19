package todo

import "github.com/elct9620/interview-gogolook-todo/internal/todo/domain"

type TaskOutput struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type CreateTaskInput struct {
	Name string `json:"name"`
}

type UpdateTaskInput struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type TodoRepository interface {
	AllTasks() []*domain.Task
	CreateTask(name string) *domain.Task
	FindTask(id int) *domain.Task
	UpdateTask(task *domain.Task) *domain.Task
	DeleteTask(id int)
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

func (uc *TodoUseCase) CreateTask(input *CreateTaskInput) *TaskOutput {
	task := uc.repo.CreateTask(input.Name)

	return &TaskOutput{
		ID:     task.ID,
		Name:   task.Name,
		Status: task.StatusCode(),
	}
}

func (uc *TodoUseCase) UpdateTask(id int, input *UpdateTaskInput) *TaskOutput {
	task := uc.repo.FindTask(id)
	task.Rename(input.Name)

	if input.Status == domain.TaskComplete {
		task.MarkComplete()
	}

	if input.Status == domain.TaskIncomplete {
		task.MarkIncomplete()
	}

	updatedTask := uc.repo.UpdateTask(task)
	return &TaskOutput{
		ID:     updatedTask.ID,
		Name:   updatedTask.Name,
		Status: updatedTask.StatusCode(),
	}
}

func (uc *TodoUseCase) DeleteTask(id int) {
	uc.repo.DeleteTask(id)
}
