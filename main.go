package main

import (
	"context"
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

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema: %v", err)
	}

	r := app.BuildRouter(client)

	r.Run()
}
