package todo

import (
	"context"
	"errors"

	"github.com/krobus00/learn-go-graphql/api/model"
	"github.com/krobus00/learn-go-graphql/api/model/database"
	"github.com/krobus00/learn-go-graphql/util"
	"github.com/microcosm-cc/bluemonday"
)

func (svc *service) Store(ctx context.Context, payload *model.CreateTodoRequest) (*model.CreateTodoResponse, error) {
	segment := util.StartTracer(ctx, tag, tracingStore)
	defer segment.End()

	p := bluemonday.UGCPolicy()

	payload = &model.CreateTodoRequest{
		Text: p.Sanitize(payload.Text),
	}

	input := &database.Todo{
		ID:     util.NewUUID(),
		Text:   payload.Text,
		IsDone: false,
	}

	err := svc.repository.TodoRepostory.Store(ctx, svc.db, input)
	if err != nil {
		return nil, err
	}

	result := &model.CreateTodoResponse{
		ID: input.ID,
	}
	return result, nil
}

func (svc *service) FindAll(ctx context.Context) ([]*model.Todo, error) {
	segment := util.StartTracer(ctx, tag, tracingFindAll)
	defer segment.End()

	results := make([]*model.Todo, 0)

	todos, err := svc.repository.TodoRepostory.FindAll(ctx, svc.db)
	if err != nil {
		return results, err
	}

	for _, v := range todos {
		results = append(results, &model.Todo{
			ID:     v.ID,
			Text:   v.Text,
			IsDone: v.IsDone,
		})
	}

	return results, nil
}

func (svc *service) Show(ctx context.Context, payload *model.GetTodoByIDRequest) (*model.Todo, error) {
	segment := util.StartTracer(ctx, tag, tracingShow)
	defer segment.End()

	input := &database.Todo{
		ID: payload.ID,
	}
	todo, err := svc.repository.TodoRepostory.FindOneByID(ctx, svc.db, input)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, errors.New("todo not found")
	}

	result := &model.Todo{
		ID:     todo.ID,
		Text:   todo.Text,
		IsDone: todo.IsDone,
	}

	return result, nil
}

func (svc *service) Update(ctx context.Context, payload *model.UpdateTodoByIDRequest) (bool, error) {
	segment := util.StartTracer(ctx, tag, tracingUpdate)
	defer segment.End()

	p := bluemonday.UGCPolicy()

	input := &database.Todo{
		ID: payload.ID,
	}

	todo, err := svc.repository.TodoRepostory.FindOneByID(ctx, svc.db, input)
	if err != nil {
		return false, err
	}
	if todo == nil {
		return false, errors.New("todo not found")
	}
	payload = &model.UpdateTodoByIDRequest{
		Text:   p.Sanitize(payload.Text),
		IsDone: payload.IsDone,
	}

	input = &database.Todo{
		ID:     todo.ID,
		Text:   payload.Text,
		IsDone: payload.IsDone,
	}
	rowsAffected, err := svc.repository.TodoRepostory.UpdateByID(ctx, svc.db, input)
	if err != nil {
		return false, err
	}
	if rowsAffected != 1 {
		return false, errors.New("update todo failed")
	}

	return true, nil
}

func (svc *service) Delete(ctx context.Context, payload *model.DeleteTodoByIDRequest) (bool, error) {
	segment := util.StartTracer(ctx, tag, tracingDelete)
	defer segment.End()

	input := &database.Todo{
		ID: payload.ID,
		BaseData: database.BaseData{
			IncludeSoftDelete: util.PointerBoolToBool(payload.IsHardDelete),
		},
	}

	todo, err := svc.repository.TodoRepostory.FindOneByID(ctx, svc.db, input)
	if err != nil {
		return false, err
	}
	if todo == nil {
		return false, errors.New("todo not found")
	}

	if util.PointerBoolToBool(payload.IsHardDelete) {
		rowsAffected, err := svc.repository.TodoRepostory.DeleteByID(ctx, svc.db, input)
		if err != nil {
			return false, err
		}
		if rowsAffected != 1 {
			return false, errors.New("delete todo failed")
		}
		return true, nil
	}

	rowsAffected, err := svc.repository.TodoRepostory.SoftDeleteByID(ctx, svc.db, input)
	if err != nil {
		return false, err
	}
	if rowsAffected != 1 {
		return false, errors.New("delete todo failed")
	}

	return true, nil
}
