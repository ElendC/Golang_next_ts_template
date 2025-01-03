package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func RegisterUser(pool *pgxpool.Pool, userName string, password string, role string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO user_table (user_name, password, role) VALUES ($1, $2, $3)
			
`
	// WithTimeout sends a cancel if more than 5 seconds passes during registration query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = pool.Exec(ctx, query, userName, string(hashedPassword), role)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("registration failed: operation timed out")
		}
		return fmt.Errorf("failed to insert user: %w", err)
	}

	fmt.Println("User successfully registered!")

	return nil
}
