package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/krobus00/learn-go-graphql/api/graph/generated"
	"github.com/krobus00/learn-go-graphql/api/model"
)

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoRequest) (*model.CreateTodoResponse, error) {
	resp, err := r.Service.TodoService.Store(ctx, &input)
	return resp, err
}

// UpdateTodoByID implements generated.MutationResolver
func (mr *mutationResolver) UpdateTodoByID(ctx context.Context, input model.UpdateTodoByIDRequest) (*model.Todo, error) {
	todo, err := mr.Service.TodoService.Update(ctx, &input)
	return todo, err
}

type queryResolver struct{ *Resolver }

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := r.Service.TodoService.FindAll(ctx)

	return todos, err
}

// TodoByID implements generated.QueryResolver
func (qr *queryResolver) TodoByID(ctx context.Context, input model.GetTodoByIDRequest) (*model.Todo, error) {
	todo, err := qr.Service.TodoService.Show(ctx, &input)
	return todo, err
}
