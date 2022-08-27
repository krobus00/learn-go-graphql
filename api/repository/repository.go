package repository

import (
	"github.com/krobus00/learn-go-graphql/api/repository/todo"
	"go.uber.org/fx"
)

type Repository struct {
	fx.In

	TodoRepostory todo.TodoRepository
}

var Module = fx.Options(
	fx.Provide(todo.New),
)
