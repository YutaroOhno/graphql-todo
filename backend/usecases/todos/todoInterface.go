package todos

import (
	"github.com/graphql-todo/backend/models"
)

type TodoUsecaseInterface interface {
	GetTodos() ([]*models.Todo, error)
	CreateTodo(input models.NewTodo) (*models.Todo, error)
}
