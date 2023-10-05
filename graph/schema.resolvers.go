package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"
	"go-enter/graph/model"
	"go-enter/internal"
)

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: CreateBook - createBook"))
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return []*model.Book{
		{
			ID:   "TODO-1",
			Title: "My Todo 1",
		},
	}, nil
}

// Mutation returns internal.MutationResolver implementation.
func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }