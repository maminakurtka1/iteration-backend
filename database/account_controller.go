package database

import (
	"context"
	"iteration-backend/dto"
	"iteration-backend/tools"

	"github.com/apex/log"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateAccount(conn *pgxpool.Pool, account *dto.AccountSignUp) (string, error) {
	defer conn.Close()
	ctx := context.Background()
	row := conn.QueryRow(ctx, "INSERT INTO accounts (id, inserted_at, updated_at) VALUES (DEFAULT, DEFAULT, DEFAULT) RETURNING id")
	var account_id string
	if err := row.Scan(&account_id); err != nil {
		log.WithError(err).Error("Can't create new account!")
		return account_id, err
	}
	password_hash, err := tools.HashPassword(account.Password)
	if err != nil {
		log.WithError(err).Error("Can't create new account!")
		return account_id, err
	}
	row = conn.QueryRow(ctx, "INSERT INTO account_identities (id, account_id, inserted_at, updated_at, phone_number, email, password_hash) VALUES (DEFAULT, $1, DEFAULT, DEFAULT, $2, $3, $4) RETURNING id", account_id, account.Phone, account.Email, password_hash)
	var id int
	if err := row.Scan(&id); err != nil {
		log.WithError(err).Error("Can't create new account!")
		return account_id, err
	}
	return account_id, err
}

func SignIn(conn *pgxpool.Pool, account *dto.AccountSignIn) (string, error) {
	defer conn.Close()
	ctx := context.Background()
	row := conn.QueryRow(ctx, "SELECT account_id, password_hash FROM account_identities WHERE phone_number=$1", account.Phone)
	var id string
	var password_hash string
	if err := row.Scan(&id, &password_hash); err != nil {
		log.WithError(err).Error("Can't find user!")
		return "", err
	}
	if tools.CheckPasswordHash(account.Password, password_hash) {
		return id, nil
	}
	return "", nil
}
