package main

import (
	"context"
	"log"

	"choseclothes/internal/app"
	"choseclothes/internal/db"
)

func main() {
	client := db.NewClient()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema: %v", err)
	}

	r := app.BuildRouter(client)

	r.Run()
}
