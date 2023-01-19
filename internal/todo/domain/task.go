package domain

const (
	TaskIncomplete = iota
	TaskComplete
)

type Task struct {
	ID        int
	Name      string
	Completed bool
}

func NewTask(id int, name string, completed bool) *Task {
	return &Task{
		ID:        id,
		Name:      name,
		Completed: completed,
	}
}

func (t *Task) StatusCode() int {
	if t.Completed {
		return TaskComplete
	}

	return TaskIncomplete
}

func (t *Task) Rename(newName string) {
	t.Name = newName
}

func (t *Task) MarkComplete() {
	t.Completed = true
}

func (t *Task) MarkIncomplete() {
	t.Completed = false
}
