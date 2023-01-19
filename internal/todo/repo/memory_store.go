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

func (s *MemoryStore) FindTask(id int) *domain.Task {
	return newTaskFromSchema(s.items[id])
}

func (s *MemoryStore) UpdateTask(task *domain.Task) *domain.Task {
	s.items[task.ID] = TaskSchema{
		ID:        task.ID,
		Name:      task.Name,
		Completed: task.Completed,
	}

	return newTaskFromSchema(s.items[task.ID])
}

func (s *MemoryStore) DeleteTask(id int) {
	_, ok := s.items[id]
	if ok {
		delete(s.items, id)
	}
}

func newTaskFromSchema(schema TaskSchema) *domain.Task {
	return domain.NewTask(schema.ID, schema.Name, schema.Completed)
}
