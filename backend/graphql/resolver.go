package graphql

import (
	"context"

	"github.com/graphql-todo/backend/models"
	"github.com/graphql-todo/backend/usecases/todos"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	TodoUsecase todos.TodoUsecaseInterface
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	return r.TodoUsecase.CreateTodo(input)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	return r.TodoUsecase.GetTodos()
}
