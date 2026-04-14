package user

type Repository interface {
	GetUsers() []User
	CreateUser(user User) error
}

type repository struct {
	users []User
}

func NewRepository() Repository {
	return &repository{
		users: []User{},
	}
}

func (r *repository) GetUsers() []User {
	return r.users
}

func (r *repository) CreateUser(user User) error {
	r.users = append(r.users, user)
	return nil
}