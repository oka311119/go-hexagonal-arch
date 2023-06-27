package main

import (
	"database/sql"
	"net/http"

	"github.com/oka311119/go-hexagonal-arch/adapter/driven"
	"github.com/oka311119/go-hexagonal-arch/adapter/driver"
	"github.com/oka311119/go-hexagonal-arch/domain/service"
)

func main() {
	db, err := sql.Open("mysql", "root:example@tcp(db:3306)/todos")
	if err != nil {
		panic(err)
	}

	repo := driven.NewMySqlTodoRepository(db)
	svc := service.NewTodoService(repo)
	handler := driver.NewHttpHandler(svc)

	http.HandleFunc("/create", handler.CreateTodoHandler)
	http.HandleFunc("/get", handler.GetTodoByIdHandler)
	http.HandleFunc("/getall", handler.GetAllTodosHandler)

	// ローカルサーバを起動
	http.ListenAndServe(":8080", nil)
}
