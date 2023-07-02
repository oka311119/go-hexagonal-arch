package driven

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/oka311119/go-hexagonal-arch/internal/domain/entity"
	"github.com/oka311119/go-hexagonal-arch/internal/domain/valueobject"
)

type MySqlTodoRepository struct {
	db *sql.DB
}

func NewMySqlTodoRepository(db *sql.DB) *MySqlTodoRepository {
	return &MySqlTodoRepository{db: db}
}

func (repo *MySqlTodoRepository) Save(todo *entity.Todo) error {
	_, err := repo.db.Exec("INSERT INTO todos (id, title, completed, duration_start, duration_end) VALUES (?, ?, ?, ?, ?)", todo.ID, todo.Title, todo.Completed, todo.Duration.Date1, todo.Duration.Date2)
	if err != nil {
		return fmt.Errorf("could not save todo to db: %v", err)
	}

	return nil
}

func (repo *MySqlTodoRepository) GetAll() ([]*entity.Todo, error) {
	rows, err := repo.db.Query("SELECT id, title, completed, duration_start, duration_end FROM todos")
	if err != nil {
		return nil, fmt.Errorf("could not get todos from db: %v", err)
	}
	defer rows.Close()

	var todos []*entity.Todo
	for rows.Next() {
		var todo entity.Todo
		var durationStart, durationEnd time.Time
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &durationStart, &durationEnd)
		if err != nil {
			return nil, fmt.Errorf("could not read row data: %v", err)
		}
		todo.Duration, err = valueobject.NewDateRange(durationStart, durationEnd)
		if err != nil {
			return nil, fmt.Errorf("could not create date range: %v", err)
		}

		todos = append(todos, &todo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate over rows: %v", err)
	}

	return todos, nil
}

func (repo *MySqlTodoRepository) GetById(id string) (*entity.Todo, error) {
	row := repo.db.QueryRow("SELECT id, title FROM todos WHERE id = ?", id)

	var todo entity.Todo
	err := row.Scan(&todo.ID, &todo.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not get todo from db: %v", err)
	}

	return &todo, nil
}
