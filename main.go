package main

import (
	"choseclothes/internal/shared/router"
	"choseclothes/internal/user"
)

func main() {
	userHandler := &user.UserHandler{}

	r := router.SetupRouter(userHandler)

	r.Run()
}
