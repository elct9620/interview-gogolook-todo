package repo

import "github.com/elct9620/interview-gogolook-todo/internal/todo/domain"

type MemoryStore struct {
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) AllTasks() []*domain.Task {
	return []*domain.Task{
		domain.NewTask(1, "name", false),
	}
}
