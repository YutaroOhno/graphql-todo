package graphql

import (
	"context"
	"todo-graphql-api/usecases/todos"
	"todo-graphql-api/entities"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
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
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*entities.Todo, error) {
	return r.TodoUsecase.GetTodos()
}
