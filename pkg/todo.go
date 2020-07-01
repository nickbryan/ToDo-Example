package pkg

import (
	"github.com/google/uuid"
)

type TodoRepository interface {
	Create(item *TodoItem) error
	Update(item *TodoItem) error
	List() TodoList
	Find(id uuid.UUID) (TodoItem, error)
}

type TodoItem struct {
	id      uuid.UUID
	content string
	done    bool
}

type TodoList []TodoItem

func NewTodoItem(content string) *TodoItem {
	return &TodoItem{
		id:      uuid.New(),
		content: content,
		done:    false,
	}
}

func (ti *TodoItem) ID() uuid.UUID {
	return ti.id
}

func (ti *TodoItem) Content() string {
	return ti.content
}

func (ti *TodoItem) Done() bool {
	return ti.done
}

func (ti *TodoItem) MarkDone(done bool) {
	ti.done = done
}

func (ti *TodoItem) UpdateContent(content string) {
	ti.content = content
}
