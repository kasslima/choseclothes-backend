package user

import "github.com/gofrs/uuid/v5"

type User struct {
	Id           uuid.UUID
	Name         string
	Email        string
	Password     string
	GoogleId     string
	IsSubscriber bool
}

type CreateUserInput struct {
	Name         string `json:"name" binding:"required,min=2,max=100"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"min=6"`
	GoogleId     string `json:"google_id"`
	IsSubscriber bool   `json:"is_subscriber"`
}
