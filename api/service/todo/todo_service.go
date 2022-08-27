package todo

import (
	"context"

	"github.com/krobus00/learn-go-graphql/api/model"
	"github.com/krobus00/learn-go-graphql/api/model/database"
	"github.com/krobus00/learn-go-graphql/util"
	"github.com/microcosm-cc/bluemonday"
)

func (svc *service) Store(ctx context.Context, payload *model.NewTodo) (*model.CreateTodoResponse, error) {
	segment := util.StartTracer(ctx, tag, tracingStore)
	defer segment.End()

	p := bluemonday.UGCPolicy()

	payload = &model.NewTodo{
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

	resp := &model.CreateTodoResponse{
		ID: input.ID,
	}
	return resp, nil
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
