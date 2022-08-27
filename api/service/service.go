package service

import (
	"github.com/krobus00/learn-go-graphql/api/service/todo"
	"go.uber.org/fx"
)

type Service struct {
	fx.In

	TodoService todo.TodoService
}

var Module = fx.Options(
	fx.Provide(todo.New),
)
