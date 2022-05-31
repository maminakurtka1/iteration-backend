package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func openConnect(ctx context.Context) *pgxpool.Pool {
	conn, err := pgxpool.Connect(ctx, "postgres://postgres:123@localhost:5432/iteration")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return conn
}
