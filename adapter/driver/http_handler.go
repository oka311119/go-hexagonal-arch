package driver

import (
	"encoding/json"
	"net/http"

	"github.com/oka311119/go-hexagonal-arch/domain/service"
)

type HttpHandler struct {
	todoService *service.TodoService
}

func NewHttpHandler(todoService *service.TodoService) *HttpHandler {
	return &HttpHandler{todoService: todoService}
}

func (h *HttpHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	// ここではリクエストからTodoを作成するためのデータを解析します
	// 実際のアプリケーションでは、エラーハンドリングとバリデーションを含むより堅牢な処理が必要になります

	var todo struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// そして、解析したデータを用いてTodoを作成します
	if err := h.todoService.Create(todo.ID, todo.Title); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HttpHandler) GetTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	// ここではリクエストからTodoのIDを取得します
	// 実際のアプリケーションでは、エラーハンドリングとバリデーションを含むより堅牢な処理が必要になります

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	// そして、取得したIDを用いてTodoを取得します
	todo, err := h.todoService.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (h *HttpHandler) GetAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}
