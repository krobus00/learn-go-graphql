package todo

import (
	"context"

	"github.com/krobus00/learn-go-graphql/api/model/database"
	"github.com/krobus00/learn-go-graphql/infrastructure"
	"go.uber.org/zap"
)

const (
	tag = `[TodoRepository]`

	tracingStore          = "Store"
	tracingFindAll        = "FindAll"
	tracingFindOneByID    = "FindOneByID"
	tracingUpdateByID     = "UpdateByID"
	tracingSoftDeleteByID = "SoftDeleteByID"
	tracingDeleteByID     = "DeleteByID"
)

type (
	TodoRepository interface {
		GetTableName() string
		Store(ctx context.Context, db infrastructure.Querier, input *database.Todo) error
		FindAll(ctx context.Context, db infrastructure.Querier) ([]*database.Todo, error)
		FindOneByID(ctx context.Context, db infrastructure.Querier, input *database.Todo) (*database.Todo, error)
		UpdateByID(ctx context.Context, db infrastructure.Querier, input *database.Todo) error
		SoftDeleteByID(ctx context.Context, db infrastructure.Querier, input *database.Todo) error
		DeleteByID(ctx context.Context, db infrastructure.Querier, input *database.Todo) error
	}
	repository struct {
		logger *zap.Logger
	}
)

func New(infrastructure infrastructure.Infrastructure) TodoRepository {
	return &repository{
		logger: infrastructure.Logger,
	}
}

func (r *repository) GetTableName() string {
	return "todos"
}
