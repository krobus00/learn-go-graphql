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
	tracingShow    = "Show"
	tracingUpdate  = "Update"
	tracingDelete  = "Delete"
)

type (
	TodoService interface {
		FindAll(ctx context.Context) ([]*model.Todo, error)
		Store(ctx context.Context, payload *model.CreateTodoRequest) (*model.CreateTodoResponse, error)
		Show(ctx context.Context, payload *model.GetTodoByIDRequest) (*model.Todo, error)
		Update(ctx context.Context, payload *model.UpdateTodoByIDRequest) (bool, error)
		Delete(ctx context.Context, payload *model.DeleteTodoByIDRequest) (bool, error)
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
