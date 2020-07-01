package rest_test

import (
	"github.com/google/uuid"
	"github.com/nickbryan/todo/pkg"
)

type mockRepository struct {
	items map[uuid.UUID]*pkg.TodoItem
}

func emptyMockRepo() *mockRepository {
	return &mockRepository{items: make(map[uuid.UUID]*pkg.TodoItem)}
}

func (m *mockRepository) Create(item *pkg.TodoItem) error {
	m.items[item.ID()] = item

	return nil
}

func (m *mockRepository) Update(item *pkg.TodoItem) error {
	m.items[item.ID()] = item

	return nil
}

func (m *mockRepository) List() pkg.TodoList {
	var list pkg.TodoList

	for _, item := range m.items {
		list = append(list, *item)
	}

	return list
}

func (m *mockRepository) Find(id uuid.UUID) (pkg.TodoItem, error) {
	var item pkg.TodoItem

	item = *m.items[id]

	return item, nil
}
