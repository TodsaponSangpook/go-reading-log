package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/todsapon/go-reading-log/config"
)

var Pool *pgxpool.Pool

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	Pool, err = pgxpool.New(ctx, config.GetDBUrl())

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
}
