package primaryport

import (
	"github.com/oka311119/go-hexagonal-arch/internal/domain/entity"
	port "github.com/oka311119/go-hexagonal-arch/internal/port/secondary"
)

type TodoUsecase interface {
	Create(id string, title string) error
	GetAll() ([]*entity.Todo, error)
	GetById(id string) (*entity.Todo, error)
}

type todoUsecase struct {
	repo port.TodoRepository
}

func NewTodoUsecase(repo port.TodoRepository) *todoUsecase {
	return &todoUsecase{repo: repo}
}

func (s *todoUsecase) Create(id string, title string) error {
	todo := &entity.Todo{ID: id, Title: title, Completed: false, Duration: nil}
	return s.repo.Save(todo)
}

func (s *todoUsecase) GetAll() ([]*entity.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoUsecase) GetById(id string) (*entity.Todo, error) {
	return s.repo.GetById(id)
}
