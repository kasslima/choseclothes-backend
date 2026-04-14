package user

import "choseclothes/ent"

type Repository interface {
	GetUsers() []User
	CreateUser(user User) error
}

type repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) GetUsers() []User {

	return nil
}

func (r *repository) CreateUser(user User) error {

	return nil
}
