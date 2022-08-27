package todo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/krobus00/learn-go-graphql/api/model"
	"github.com/krobus00/learn-go-graphql/api/repository"
	"github.com/krobus00/learn-go-graphql/infrastructure"

	"go.uber.org/zap"
)

const (
	tag = `[TodoService]`

	tracingStore   = "Store"
	tracingFindAll = "FindAll"
)

type (
	TodoService interface {
		FindAll(ctx context.Context) ([]*model.Todo, error)
		Store(ctx context.Context, payload *model.NewTodo) (*model.CreateTodoResponse, error)
	}
	service struct {
		logger *zap.Logger
		config *infrastructure.Config
		db     *sqlx.DB

		repository repository.Repository
	}
)

func New(
	infrastructure infrastructure.Infrastructure,
	repository repository.Repository,
) TodoService {
	return &service{
		logger:     infrastructure.Logger,
		config:     infrastructure.Config,
		db:         infrastructure.Database.SqlxDB,
		repository: repository,
	}
}
