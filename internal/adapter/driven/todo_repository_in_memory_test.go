package driven

import (
	"testing"
	"time"

	"github.com/oka311119/go-hexagonal-arch/internal/domain/entity"
	"github.com/oka311119/go-hexagonal-arch/internal/domain/valueobject"
)

func TestInMemoryTodoRepository(t *testing.T) {
	repo := NewInMemoryTodoRepository()

	todo := &entity.Todo{
		ID:        "1",
		Title:     "test todo",
		Completed: false,
		Duration: &valueobject.DateRange{
			Date1: time.Date(2023, 7, 2, 0, 0, 0, 0, time.UTC),
			Date2: time.Date(2023, 7, 4, 0, 0, 0, 0, time.UTC),
		},
	}

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
