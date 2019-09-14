package todos

import (
	"todo-graphql-api/entities"
)

type TodoUsecaseInterface interface {
	GetTodos()([]*entities.Todo, error)
}
