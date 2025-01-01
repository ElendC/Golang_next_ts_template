package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func DbConfig() {

	// Connect to the database using the DATABASE_URL from the environment
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Testing connection
	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world. Database connection succeed!'").Scan(&greeting) //Scan maps query results into variables
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(greeting)
}
