package driven

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/oka311119/go-hexagonal-arch/internal/domain/entity"
)

func TestSaveAndGetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create the mock repository
	mockRepo := NewMockTodoRepository(ctrl)

	// Define the todo to be used in tests
	todo := &entity.Todo{ID: "1", Title: "My first task"}

	// Expectations
	mockRepo.EXPECT().Save(todo).Return(nil)
	mockRepo.EXPECT().GetById("1").Return(todo, nil)

	err := mockRepo.Save(todo)
	if err != nil {
		t.Fatalf("could not save todo: %v", err)
	}

	retrievedTodo, err := mockRepo.GetById("1")
	if err != nil {
		t.Fatalf("could not get todo: %v", err)
	}

	if retrievedTodo.ID != todo.ID || retrievedTodo.Title != todo.Title {
		t.Errorf("retrieved todo does not match saved todo")
	}
}
