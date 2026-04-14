package user

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

type Service interface {
	GetUsers() []User
	CreateUser(user User) (User, error)
}

func (s *service) GetUsers() []User {
	return s.repo.GetUsers()
}

func (s *service) CreateUser(user User) (User, error) {


	err := s.repo.CreateUser(user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}