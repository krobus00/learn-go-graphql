package todo

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/krobus00/learn-go-graphql/api/model/database"
)

func (r *repository) buildInsertQuery(input *database.Todo) sq.InsertBuilder {
	vals := sq.Eq{
		"id":         input.ID,
		"text":       input.Text,
		"is_done":    input.IsDone,
		"created_at": time.Now().Unix(),
		"updated_at": time.Now().Unix(),
	}
	insertBuilder := sq.Insert(r.GetTableName()).SetMap(vals)
	return insertBuilder
}

func (r *repository) buildUpdateQuery(input *database.Todo) sq.UpdateBuilder {
	vals := sq.Eq{
		"text":       input.Text,
		"is_done":    input.IsDone,
		"updated_at": time.Now().Unix(),
	}
	updateBuilder := sq.Update(r.GetTableName()).SetMap(vals)
	return updateBuilder
}

func (r *repository) buildSoftDeleteQuery(input *database.Todo) sq.UpdateBuilder {
	vals := sq.Eq{
		"updated_at": time.Now().Unix(),
		"deleted_at": time.Now().Unix(),
	}
	updateBuilder := sq.Update(r.GetTableName()).SetMap(vals)
	return updateBuilder
}

func (r *repository) buildDeleteQuery(input *database.Todo) sq.DeleteBuilder {
	deleteBuilder := sq.Delete(r.GetTableName())
	return deleteBuilder
}

func (r *repository) buildSelectQuery(input *database.Todo) sq.SelectBuilder {
	selection := []string{
		"id",
		"text",
		"is_done",
		"created_at",
		"updated_at",
		"deleted_at",
	}
	selectBuilder := sq.Select(selection...).From(r.GetTableName())
	if input != nil && !input.IncludeSoftDelete {
		selectBuilder = selectBuilder.Where(sq.Eq{"deleted_at": nil})
	}
	return selectBuilder
}
