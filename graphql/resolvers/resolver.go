package resolvers

import (
	"context"

	"github.com/jukemori/gqlgen-todos/infra/mysql"
)

type TodoDB interface {
    GetUserByID(ctx context.Context, id string) (*mysql.User, error)
    GetAllTodos(ctx context.Context) ([]*mysql.Todo, error)
    CreateTodo(ctx context.Context, todo *mysql.Todo) error
}

type Resolver struct {
    DB TodoDB
}