package repo

import "github.com/elct9620/interview-gogolook-todo/internal/todo/domain"

type TaskSchema struct {
	ID        int
	Name      string
	Completed bool
}

type MemoryStore struct {
	pointer int
	items   map[int]TaskSchema
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		pointer: 0,
		items:   make(map[int]TaskSchema),
	}
}

func (s *MemoryStore) NextID() int {
	s.pointer += 1
	return s.pointer
}

func (s *MemoryStore) AllTasks() []*domain.Task {
	tasks := []*domain.Task{}

	for _, item := range s.items {
		tasks = append(tasks, newTaskFromSchema(item))
	}

	return tasks
}

func (s *MemoryStore) CreateTask(name string) *domain.Task {
	nextID := s.NextID()
	s.items[nextID] = TaskSchema{
		ID:        nextID,
		Name:      name,
		Completed: false,
	}

	return newTaskFromSchema(s.items[nextID])
}

func newTaskFromSchema(schema TaskSchema) *domain.Task {
	return domain.NewTask(schema.ID, schema.Name, schema.Completed)
}
