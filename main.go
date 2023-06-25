package main

import (
	"net/http"

	"github.com/oka311119/go-hexagonal-arch/adapter/driven"
	"github.com/oka311119/go-hexagonal-arch/adapter/driver"
	"github.com/oka311119/go-hexagonal-arch/domain/service"
)

func main() {
	// In-memory Todoリポジトリを作成
	repo := driven.NewInMemoryTodoRepository()

	// Todoサービスを作成
	todoService := service.NewTodoService(repo)

	// HTTPハンドラを作成
	httpHandler := driver.NewHttpHandler(todoService)

	// 各エンドポイントにハンドラを割り当て
	http.HandleFunc("/create", httpHandler.CreateTodoHandler)
	http.HandleFunc("/get", httpHandler.GetTodoByIdHandler)
	http.HandleFunc("/getall", httpHandler.GetAllTodosHandler)

	// ローカルサーバを起動
	http.ListenAndServe(":8080", nil)
}
