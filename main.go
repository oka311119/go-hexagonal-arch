package main

import (
	"net/http"

	"github.com/oka311119/go-hexagonal-arch/adapter/driven"
	"github.com/oka311119/go-hexagonal-arch/adapter/driver"
	"github.com/oka311119/go-hexagonal-arch/domain/service"
)

func main() {
	repo := driven.NewInMemoryTodoRepository()
	todoService := service.NewTodoService(repo)
	httpHandler := driver.NewHttpHandler(todoService)

	http.HandleFunc("/create", httpHandler.CreateTodoHandler)
	http.HandleFunc("/get", httpHandler.GetTodoByIdHandler)
	http.HandleFunc("/getall", httpHandler.GetAllTodosHandler)

	// ローカルサーバを起動
	http.ListenAndServe(":8080", nil)
}
