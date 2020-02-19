// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
package graphql

import (
	"context"

	"github.com/graphql-todo/backend/models"
	"github.com/graphql-todo/backend/infrastructure/todos"
)

type Resolver struct{
	// 一旦直でusecase呼び出し。
	TodoUsecase	*todos.TodoUsecase
}


func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	return r.TodoUsecase.CreateTodo(input)
	// panic("not implemented")
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	// panic("not implemented")
	return r.TodoUsecase.GetTodos()
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
