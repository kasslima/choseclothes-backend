package user

import (
	"context"

	"choseclothes/ent"
)

type Repository interface {
	GetUsers(ctx context.Context) ([]*ent.User, error)
	CreateUser(ctx context.Context, u CreateUserInput) (*ent.User, error)
}

type repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) GetUsers(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

func (r *repository) CreateUser(ctx context.Context, u CreateUserInput) (*ent.User, error) {
	return r.client.User.
		Create().
		SetNillableGoogleID(nillableString(u.GoogleId)).
		SetEmail(u.Email).
		SetNillablePasswordHash(nillableString(u.Password)).
		SetNillableName(nillableString(u.Name)).
		SetIsSubscriber(u.IsSubscriber).
		Save(ctx)
}

// helper simples pra optional string
func nillableString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}