package service

import (
	"github.com/oka311119/go-hexagonal-arch/domain/entity"
	"github.com/oka311119/go-hexagonal-arch/domain/port"
)

type TodoService struct {
	repo port.TodoRepository
}

func NewTodoService(repo port.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) Create(id string, title string) error {
	todo := &entity.Todo{ID: id, Title: title}
	return s.repo.Save(todo)
}

func (s *TodoService) GetAll() ([]*entity.Todo, error) {
	return s.repo.GetAll()
}

func (s *TodoService) GetById(id string) (*entity.Todo, error) {
	return s.repo.GetById(id)
}
