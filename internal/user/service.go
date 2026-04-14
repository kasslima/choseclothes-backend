package user

import (
	"context"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

type Service interface {
	GetUsers(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, input CreateUserInput) (User, error)
}

func (s *service) GetUsers(ctx context.Context) ([]User, error) {
	users, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]User, 0, len(users))

	for _, u := range users {
		result = append(result, User{
			Id:           u.ID,
			Name:         u.Name,
			Email:        u.Email,
			GoogleId:     u.GoogleID,
			IsSubscriber: u.IsSubscriber,
		})
	}

	return result, nil
}

func (s *service) CreateUser(ctx context.Context, input CreateUserInput) (User, error) {
	u, err := s.repo.CreateUser(ctx, input)
	if err != nil {
		return User{}, err
	}

	return User{
		Id:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		GoogleId:     u.GoogleID,
		IsSubscriber: u.IsSubscriber,
	}, nil
}
