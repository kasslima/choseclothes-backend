package db

import (
	"log"
	"os"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	_ "github.com/lib/pq"

	"choseclothes/ent"
)

func NewClient() *ent.Client {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	drv, err := sql.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatalf("failed opening connection: %v", err)
	}

	return ent.NewClient(ent.Driver(drv))
}
