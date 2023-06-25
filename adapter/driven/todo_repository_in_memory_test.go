package driven

import (
	"testing"

	"github.com/oka311119/go-hexagonal-arch/domain/entity"
)

func TestInMemoryTodoRepository(t *testing.T) {
	repo := NewInMemoryTodoRepository()

	todo := &entity.Todo{ID: "1", Title: "test todo"}

	err := repo.Save(todo)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	gotTodo, err := repo.GetById("1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if gotTodo.ID != todo.ID || gotTodo.Title != todo.Title {
		t.Errorf("Got todo does not match the saved one: %v", gotTodo)
	}
}
