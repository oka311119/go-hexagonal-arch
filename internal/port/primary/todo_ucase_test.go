package primaryport

import (
	"testing"

	driven "github.com/oka311119/go-hexagonal-arch/internal/adapter/driven"
)

func TestTodoUsecase(t *testing.T) {
	repo := driven.NewInMemoryTodoRepository()

	todoUsecase := NewTodoUsecase(repo)

	err := todoUsecase.Create("1", "test todo")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	todo, err := todoUsecase.GetById("1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if todo.ID != "1" || todo.Title != "test todo" {
		t.Errorf("Todo does not match expected: %v", todo)
	}
}
