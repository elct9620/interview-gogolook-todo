package repo_test

import (
	"testing"

	"github.com/elct9620/interview-gogolook-todo/internal/todo/repo"
	"github.com/google/go-cmp/cmp"
)

func Test_MemoryStore_AllTasks(t *testing.T) {
	store := repo.NewMemoryStore()
	store.CreateTask("買晚餐")
	store.CreateTask("買早餐")

	tasks := store.AllTasks()
	if !cmp.Equal(2, len(tasks)) {
		t.Fail()
	}
}

func Test_MemoryStore_FindTask(t *testing.T) {
	store := repo.NewMemoryStore()
	store.CreateTask("買晚餐")
	store.CreateTask("買早餐")

	task := store.FindTask(1)
	if !cmp.Equal("買晚餐", task.Name) {
		t.Error(cmp.Diff("買晚餐", task.Name))
	}

	task = store.FindTask(2)
	if !cmp.Equal("買早餐", task.Name) {
		t.Error(cmp.Diff("買早餐", task.Name))
	}

	task = store.FindTask(3)
	if task != nil {
		t.Fail()
	}
}

func Test_MemoryStore_UpdateTask(t *testing.T) {
	store := repo.NewMemoryStore()
	store.CreateTask("買晚餐")
	store.CreateTask("買早餐")

	task := store.FindTask(1)
	task.Rename("買午餐")
	store.UpdateTask(task)

	updatedTask := store.FindTask(1)
	if !cmp.Equal("買午餐", updatedTask.Name) {
		t.Error(cmp.Diff("買午餐", updatedTask.Name))
	}
}

func Test_MemoryStore_DeleteTask(t *testing.T) {
	store := repo.NewMemoryStore()
	store.CreateTask("買晚餐")
	store.CreateTask("買早餐")

	store.DeleteTask(1)
	tasks := store.AllTasks()

	if !cmp.Equal(1, len(tasks)) {
		t.Fail()
	}
}
