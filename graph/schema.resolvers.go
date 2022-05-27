package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"prevention_productivity/graph/generated"
	"prevention_productivity/graph/model"
)

func (r *grantResolver) TeamMembers(ctx context.Context, obj *model.Grant) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddUser(ctx context.Context, input *model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *noteResolver) ID(ctx context.Context, obj *model.Note) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productivityLogResolver) ApprovalStatus(ctx context.Context, obj *model.ProductivityLog) (model.ApprovalStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productivityLogResolver) Notes(ctx context.Context, obj *model.ProductivityLog) ([]*model.Note, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Grant returns generated.GrantResolver implementation.
func (r *Resolver) Grant() generated.GrantResolver { return &grantResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Note returns generated.NoteResolver implementation.
func (r *Resolver) Note() generated.NoteResolver { return &noteResolver{r} }

// ProductivityLog returns generated.ProductivityLogResolver implementation.
func (r *Resolver) ProductivityLog() generated.ProductivityLogResolver {
	return &productivityLogResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type grantResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type noteResolver struct{ *Resolver }
type productivityLogResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
