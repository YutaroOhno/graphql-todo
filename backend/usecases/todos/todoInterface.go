package todos

import (
	"github.com/graphql-todo/backend/entities"
)

type TodoUsecaseInterface interface {
	GetTodos() ([]*entities.Todo, error)
	CreateTodo(input entities.NewTodo) (*entities.Todo, error)
}
