package database

import (
	"context"
	_ "github.com/lib/pq"
	"go-boilerplate/infrastructure/ent"
	"go.uber.org/zap"
	"os"
)

func InitDb(sugar *zap.SugaredLogger) *ent.Client {
	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		sugar.Fatalf("failed opening connection to postgres: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		sugar.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
