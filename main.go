package main

import (
	"log"

	"github.com/joho/godotenv"

	"choseclothes/internal/app"
	"choseclothes/internal/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("warning: could not load .env file: %v", err)
	}

	client := db.NewClient()

	r := app.BuildRouter(client)

	r.Run()
}
