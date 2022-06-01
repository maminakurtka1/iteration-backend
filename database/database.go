package database

import (
	"context"

	"github.com/apex/log"
	"github.com/jackc/pgx/v4/pgxpool"
)

func openConnect(ctx context.Context) *pgxpool.Pool {
	conn, err := pgxpool.Connect(ctx, "postgres://postgres:123@localhost:5432/iteration")
	if err != nil {
		log.WithError(err).Error("Error with open database")
	}
	return conn
}
