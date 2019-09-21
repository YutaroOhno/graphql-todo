package graphql

import (
	"context"
	"graphql-todo/backend/entities"
	"graphql-todo/backend/usecases/todos"
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

func (r *mutationResolver) CreateTodo(ctx context.Context, input entities.NewTodo) (*entities.Todo, error) {
	return r.TodoUsecase.CreateTodo(input)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*entities.Todo, error) {
	return r.TodoUsecase.GetTodos()
}
