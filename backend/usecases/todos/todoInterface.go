package todos

import (
	"graphql-todo/backend/entities"
)

type TodoUsecaseInterface interface {
	GetTodos() ([]*entities.Todo, error)
	CreateTodo(input *entities.NewTodo) (*entities.Todo, error)
}
