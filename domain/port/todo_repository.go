package port

import "github.com/oka311119/go-hexagonal-arch/domain/entity"

type TodoRepository interface {
	Save(todo *entity.Todo) error
	GetAll() ([]*entity.Todo, error)
	GetById(id string) (*entity.Todo, error)
}
