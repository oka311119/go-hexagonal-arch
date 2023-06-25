package driver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oka311119/go-hexagonal-arch/adapter/driven"
	"github.com/oka311119/go-hexagonal-arch/domain/service"
)

func TestHttpHandler(t *testing.T) {
	repo := driven.NewInMemoryTodoRepository()
	svc := service.NewTodoService(repo)
	handler := NewHttpHandler(svc)

	todoJSON := `{"id":"1","title":"test todo"}`

	// Test CreateTodoHandler
	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer([]byte(todoJSON)))
	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler.CreateTodoHandler)

	handlerFunc.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Test GetTodoByIdHandler
	req, _ = http.NewRequest("GET", "/getbyid?id=1", nil)
	rr = httptest.NewRecorder()
	handlerFunc = http.HandlerFunc(handler.GetTodoByIdHandler)

	handlerFunc.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var todo map[string]string
	json.Unmarshal(rr.Body.Bytes(), &todo)
	if todo["ID"] != "1" || todo["Title"] != "test todo" {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}

	// Test GetAllTodosHandler
	req, _ = http.NewRequest("GET", "/getall", nil)
	rr = httptest.NewRecorder()
	handlerFunc = http.HandlerFunc(handler.GetAllTodosHandler)

	handlerFunc.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var todos []map[string]string
	json.Unmarshal(rr.Body.Bytes(), &todos)

	if len(todos) != 1 || todos[0]["ID"] != "1" || todos[0]["Title"] != "test todo" {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}
