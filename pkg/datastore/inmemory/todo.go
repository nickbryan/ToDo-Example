package inmemory

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nickbryan/todo/pkg"
)

type TodoRepository struct {
	items map[uuid.UUID]*pkg.TodoItem
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		items: make(map[uuid.UUID]*pkg.TodoItem),
	}
}

func (tr *TodoRepository) Create(item *pkg.TodoItem) error {
	if _, ok := tr.items[item.ID()]; ok {
		return errors.New(fmt.Sprintf("item with id %s already exists", item.ID()))
	}

	tr.items[item.ID()] = item

	return nil
}

func (tr *TodoRepository) Update(item *pkg.TodoItem) error {
	if _, ok := tr.items[item.ID()]; !ok {
		return errors.New(fmt.Sprintf("item with id %s does not exist", item.ID()))
	}

	tr.items[item.ID()] = item

	return nil
}

func (tr *TodoRepository) List() pkg.TodoList {
	var list pkg.TodoList

	for _, item := range tr.items {
		list = append(list, *item)
	}

	return list
}

func (tr *TodoRepository) Find(id uuid.UUID) (pkg.TodoItem, error) {
	var item pkg.TodoItem

	if _, ok := tr.items[id]; !ok {
		return item, errors.New(fmt.Sprintf("item with id %s does not exist", id))
	}

	item = *tr.items[id]

	return item, nil
}
