package user

type User struct {
	Id           int8   `json:"id" binding:"required,gte=1"`
	Name         string `json:"name" binding:"required,min=2,max=100"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"min=6"`
	GoogleId     string `json:"google_id"`
	IsSubscriber bool   `json:"is_subscriber"`
}
