package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUserTable(pool *pgxpool.Pool) error {
	query := `
		CREATE TABLE IF NOT EXISTS user_table (
			id UUID PRIMARY KEY,
			name TEXT NOT NULL,
			pass_hash TEXT NOT NULL,
			role TEXT, 
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err := pool.Exec(context.Background(), query)

	if err != nil {
		return fmt.Errorf("failed to create user table: %w", err)
	}

	fmt.Println("User table created successfully!")
	return nil
}
