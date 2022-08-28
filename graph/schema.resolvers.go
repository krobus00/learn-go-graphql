package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/krobus00/learn-go-graphql/api/graph/generated"
	"github.com/krobus00/learn-go-graphql/api/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoRequest) (*model.CreateTodoResponse, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// UpdateTodoByID is the resolver for the updateTodoByID field.
func (r *mutationResolver) UpdateTodoByID(ctx context.Context, input model.UpdateTodoByIDRequest) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: UpdateTodoByID - updateTodoByID"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// TodoByID is the resolver for the todoByID field.
func (r *queryResolver) TodoByID(ctx context.Context, input model.GetTodoByIDRequest) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: TodoByID - todoByID"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
