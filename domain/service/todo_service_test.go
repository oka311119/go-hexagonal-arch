package service

import (
	"testing"

	"github.com/oka311119/go-hexagonal-arch/adapter/driven"
)

func TestTodoService(t *testing.T) {
	// In-memory Todoリポジトリを作成
	repo := driven.NewInMemoryTodoRepository()

	// Todoサービスを作成
	todoService := NewTodoService(repo)

	// Todoを作成
	err := todoService.Create("1", "test todo")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Todoを取得
	todo, err := todoService.GetById("1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Todoの内容を確認
	if todo.ID != "1" || todo.Title != "test todo" {
		t.Errorf("Todo does not match expected: %v", todo)
	}
}
