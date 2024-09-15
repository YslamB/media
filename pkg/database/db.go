package database

import (
	"context"
	"fmt"
	"media/pkg/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// InitDB initializes the database connection pool.
func InitDB() *pgxpool.Pool {
	connectionString := buildConnectionString()

	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		panic(fmt.Sprintf("failed to create connection pool: %v", err))
	}

	DB = pool

	return pool
}

// buildConnectionString constructs the connection string for the database.
func buildConnectionString() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		config.ENV.DB_USER, config.ENV.DB_PASSWORD,
		config.ENV.DB_HOST, config.ENV.DB_PORT, config.ENV.DB_NAME,
	)
}
