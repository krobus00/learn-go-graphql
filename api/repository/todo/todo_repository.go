package todo

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/krobus00/learn-go-graphql/api/model/database"
	"github.com/krobus00/learn-go-graphql/infrastructure"
	"github.com/krobus00/learn-go-graphql/util"
)

func (r *repository) Store(ctx context.Context, db infrastructure.Querier, input *database.Todo) error {
	segment := util.StartTracer(ctx, tag, tracingStore)
	defer segment.End()

	query, args, err := r.buildInsertQuery(input).ToSql()
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingStore, err)
		return err
	}
	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingStore, err)
		return err
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context, db infrastructure.Querier) ([]*database.Todo, error) {
	segment := util.StartTracer(ctx, tag, tracingFindAll)
	defer segment.End()

	results := make([]*database.Todo, 0)

	query, args, err := r.buildSelectQuery().ToSql()
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingFindAll, err)
		return results, err
	}

	err = db.SelectContext(ctx, &results, query, args...)
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingFindAll, err)
		return results, err
	}

	return results, nil
}

func (r *repository) FindOneByID(ctx context.Context, db infrastructure.Querier, input *database.Todo) (*database.Todo, error) {
	segment := util.StartTracer(ctx, tag, tracingFindOneByID)
	defer segment.End()

	result := new(database.Todo)

	query, args, err := r.buildSelectQuery().Where(squirrel.Eq{
		"id": input.ID,
	}).ToSql()
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingFindOneByID, err)
		return result, err
	}

	err = db.GetContext(ctx, result, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		util.ErrorLogger(r.logger, tag, tracingFindOneByID, err)
		return result, err
	}

	return result, nil
}

func (r *repository) UpdateByID(ctx context.Context, db infrastructure.Querier, input *database.Todo) error {
	segment := util.StartTracer(ctx, tag, tracingUpdateByID)
	defer segment.End()

	query, args, err := r.buildUpdateQuery(input).Where(squirrel.Eq{
		"id": input.ID,
	}).ToSql()
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingUpdateByID, err)
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingUpdateByID, err)
		return err
	}

	return nil
}
