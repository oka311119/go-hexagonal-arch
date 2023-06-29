package driver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oka311119/go-hexagonal-arch/internal/adapter/driven"
	usecase "github.com/oka311119/go-hexagonal-arch/internal/port/primary"
)

// 共通の初期化処理
func setup() (*HttpHandler, *httptest.ResponseRecorder) {
	repo := driven.NewInMemoryTodoRepository()
	uc := usecase.NewTodoUsecase(repo)
	handler := NewHttpHandler(uc)
	rr := httptest.NewRecorder()
	return handler, rr
}

// CreateTodoHandlerテストとGet系テストで共通するTodo作成処理
func createTestTodo(handler *HttpHandler, rr *httptest.ResponseRecorder) {
	todoJSON := `{"id":"1","title":"test todo"}`
	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer([]byte(todoJSON)))
	handlerFunc := http.HandlerFunc(handler.CreateTodoHandler)
	handlerFunc.ServeHTTP(rr, req)
}

func TestCreateTodoHandler(t *testing.T) {
	handler, rr := setup()
	createTestTodo(handler, rr)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestGetAllTodosHandler(t *testing.T) {
	handler, rr := setup()
	createTestTodo(handler, rr)

	req, _ := http.NewRequest("GET", "/getall", nil)
	rr = httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler.GetAllTodosHandler)

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

func TestGetTodoByIdHandler(t *testing.T) {
	handler, rr := setup()
	createTestTodo(handler, rr)

	req, _ := http.NewRequest("GET", "/getbyid?id=1", nil)
	rr = httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler.GetTodoByIdHandler)

	handlerFunc.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var todo map[string]string
	json.Unmarshal(rr.Body.Bytes(), &todo)
	if todo["ID"] != "1" || todo["Title"] != "test todo" {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}
