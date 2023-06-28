package driver

import (
	"encoding/json"
	"net/http"

	service "github.com/oka311119/go-hexagonal-arch/internal/port/primary"
)

type HttpHandler struct {
	todoService *service.TodoService
}

func NewHttpHandler(todoService *service.TodoService) *HttpHandler {
	return &HttpHandler{todoService: todoService}
}

func (h *HttpHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.todoService.Create(todo.ID, todo.Title); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HttpHandler) GetAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func (h *HttpHandler) GetTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	todo, err := h.todoService.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}
