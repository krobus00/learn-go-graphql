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

	query, args, err := r.buildSelectQuery(nil).ToSql()
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

	query, args, err := r.buildSelectQuery(input).Where(squirrel.Eq{
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

func (r *repository) UpdateByID(ctx context.Context, db infrastructure.Querier, input *database.Todo) (int64, error) {
	segment := util.StartTracer(ctx, tag, tracingUpdateByID)
	defer segment.End()

	query, args, err := r.buildUpdateQuery(input).Where(squirrel.Eq{
		"id": input.ID,
	}).ToSql()
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingUpdateByID, err)
		return 0, err
	}

	row, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingUpdateByID, err)
		return 0, err
	}
	rowsAffected, err := row.RowsAffected()
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingUpdateByID, err)
		return 0, err
	}

	return rowsAffected, nil
}

func (r *repository) SoftDeleteByID(ctx context.Context, db infrastructure.Querier, input *database.Todo) (int64, error) {
	segment := util.StartTracer(ctx, tag, tracingSoftDeleteByID)
	defer segment.End()

	query, args, err := r.buildSoftDeleteQuery(input).Where(squirrel.Eq{
		"id": input.ID,
	}).ToSql()
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingSoftDeleteByID, err)
		return 0, err
	}

	row, err := db.ExecContext(ctx, query, args...)

	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingSoftDeleteByID, err)
		return 0, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingUpdateByID, err)
		return 0, err
	}

	return rowsAffected, nil
}

func (r *repository) DeleteByID(ctx context.Context, db infrastructure.Querier, input *database.Todo) (int64, error) {
	segment := util.StartTracer(ctx, tag, tracingDeleteByID)
	defer segment.End()

	query, args, err := r.buildDeleteQuery(input).Where(squirrel.Eq{
		"id": input.ID,
	}).ToSql()
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingDeleteByID, err)
		return 0, err
	}

	row, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingDeleteByID, err)
		return 0, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		util.ErrorLogger(r.logger, tag, tracingUpdateByID, err)
		return 0, err
	}

	return rowsAffected, nil
}
