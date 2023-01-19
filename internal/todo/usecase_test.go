package todo_test

import (
	"testing"

	"github.com/elct9620/interview-gogolook-todo/internal/todo"
	"github.com/elct9620/interview-gogolook-todo/internal/todo/repo"
	"github.com/google/go-cmp/cmp"
)

func Test_UseCase_ListTasks(t *testing.T) {
	store := repo.NewMemoryStore()
	store.CreateTask("買晚餐")
	store.CreateTask("買早餐")

	usecase := todo.NewTodoUseCase(store)
	tasks := usecase.ListTasks()

	if !cmp.Equal(2, len(tasks)) {
		t.Fail()
	}
}

func Test_UseCase_CreateTask(t *testing.T) {
	store := repo.NewMemoryStore()
	usecase := todo.NewTodoUseCase(store)

	task := usecase.CreateTask(&todo.CreateTaskInput{Name: "買早餐"})

	if !cmp.Equal("買早餐", task.Name) {
		t.Error(cmp.Diff("買早餐", task.Name))
	}
}

func Test_UseCase_UpdateTask(t *testing.T) {
	store := repo.NewMemoryStore()
	store.CreateTask("買晚餐")

	usecase := todo.NewTodoUseCase(store)
	task := usecase.UpdateTask(1, &todo.UpdateTaskInput{1, "買午餐", 0})

	if !cmp.Equal("買午餐", task.Name) {
		t.Error(cmp.Diff("買午餐", task.Name))
	}
}

func Test_UseCase_DeleteTask(t *testing.T) {
	store := repo.NewMemoryStore()
	store.CreateTask("買晚餐")

	usecase := todo.NewTodoUseCase(store)
	usecase.DeleteTask(1)

	tasks := usecase.ListTasks()
	if !cmp.Equal(0, len(tasks)) {
		t.Fail()
	}
}
