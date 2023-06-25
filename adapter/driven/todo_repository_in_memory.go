package driven

import (
	"errors"

	"github.com/oka311119/go-hexagonal-arch/domain/entity"
	"github.com/oka311119/go-hexagonal-arch/domain/port"
)

type InMemoryTodoRepository struct {
	store map[string]*entity.Todo
}

func NewInMemoryTodoRepository() port.TodoRepository {
	return &InMemoryTodoRepository{store: make(map[string]*entity.Todo)}
}

func (r *InMemoryTodoRepository) Save(todo *entity.Todo) error {
	r.store[todo.ID] = todo
	return nil
}

func (r *InMemoryTodoRepository) GetAll() ([]*entity.Todo, error) {
	var todos []*entity.Todo
	for _, todo := range r.store {
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *InMemoryTodoRepository) GetById(id string) (*entity.Todo, error) {
	todo, exists := r.store[id]
	if !exists {
		return nil, errors.New("todo not found")
	}
	return todo, nil
}
