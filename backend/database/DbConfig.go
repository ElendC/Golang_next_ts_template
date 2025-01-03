package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

// App will hold the instance for pgx.pool.Pool which is the database connection.
type App struct {
	DB *pgxpool.Pool
}

// DbConfig sets up a pool connection to the database and create tables
func DbConfig() (*pgxpool.Pool, error) {

	// Initialize connection pool
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	app := &App{DB: pool}

	// Creating tables
	err2 := CreateUserTable(app.DB)
	if err2 != nil {
		return nil, fmt.Errorf("unable to create user table: %v", err2)
	}

	return pool, nil
}
