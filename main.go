package main

import (
	"choseclothes/internal/http"
	"choseclothes/internal/user"
)

func main() {
	userHandler := &user.UserHandler{}

	r := http.SetupRouter(userHandler)

	r.Run()
}
