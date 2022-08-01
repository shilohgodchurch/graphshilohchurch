package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"shilohchurch/graph/generated"
	"shilohchurch/graph/model"
)

// CreateMovie is the resolver for the createMovie field.
// schema.resolvers.go

func (r *mutationResolver) CreateMovie(ctx context.Context, input model.NewMovie) (*model.Movie, error) {
	movie := model.Movie{
		Title: input.Title,
		URL:   input.URL,
	}

	_, err := r.DB.Model(&movie).Insert()
	if err != nil {
		return nil, fmt.Errorf("error inserting new movie: %v", err)
	}

	return &movie, nil
}

// Movies is the resolver for the movies field.
// schema.resolvers.go

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	var movies []*model.Movie

	err := r.DB.Model(&movies).Select()
	if err != nil {
		return nil, err
	}

	return movies, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
