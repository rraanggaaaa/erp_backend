package config

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDatabase() error {

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		return errors.New("DATABASE_URL is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return err
	}

	if err := pool.Ping(ctx); err != nil {
		return err
	}

	DB = pool

	fmt.Println("✅ PostgreSQL Connected")

	return nil
}
