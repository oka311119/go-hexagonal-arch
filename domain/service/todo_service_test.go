package service

import (
	"testing"

	"github.com/oka311119/go-hexagonal-arch/adapter/driven"
)

func TestTodoService(t *testing.T) {
	repo := driven.NewInMemoryTodoRepository()

	todoService := NewTodoService(repo)

	err := todoService.Create("1", "test todo")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	todo, err := todoService.GetById("1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if todo.ID != "1" || todo.Title != "test todo" {
		t.Errorf("Todo does not match expected: %v", todo)
	}
}
