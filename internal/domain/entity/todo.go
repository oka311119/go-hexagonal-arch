package entity

import "github.com/oka311119/go-hexagonal-arch/internal/domain/valueobject"

type Todo struct {
	ID        string
	Title     string
	Completed bool
	Duration  *valueobject.DateRange
}
