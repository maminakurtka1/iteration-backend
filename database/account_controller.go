package database

import (
	"context"
	"fmt"
	"iteration-backend/dto"
	"iteration-backend/tools"
	"os"
)

func CreateAccount(account *dto.AccountSignUp) (string, error) {
	ctx := context.Background()
	conn := openConnect(ctx)
	defer conn.Close()
	row := conn.QueryRow(ctx, "INSERT INTO accounts (id, inserted_at, updated_at) VALUES (DEFAULT, DEFAULT, DEFAULT) RETURNING id")
	var account_id string
	if err := row.Scan(&account_id); err != nil {
		fmt.Fprintf(os.Stderr, "Can't create new account: %v\n", err)
	}
	password_hash, err := tools.HashPassword(account.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't hash password: %v\n", err)
	}
	row = conn.QueryRow(ctx, "INSERT INTO account_identities (id, account_id, inserted_at, updated_at, phone_number, email, password_hash) VALUES (DEFAULT, $1, DEFAULT, DEFAULT, $2, $3, $4) RETURNING id", account_id, account.Phone, account.Email, password_hash)
	var id int
	if err := row.Scan(&id); err != nil {
		fmt.Fprintf(os.Stderr, "Can't create new account: %v\n", err)
	}
	return account_id, err
}

func SignIn(account *dto.AccountSignIn) (string, error) {
	ctx := context.Background()
	conn := openConnect(ctx)
	defer conn.Close()

	row := conn.QueryRow(ctx, "SELECT account_id, password_hash FROM account_identities WHERE phone_number=$1", account.Phone)
	var id string
	var password_hash string
	if err := row.Scan(&id, &password_hash); err != nil {
		fmt.Fprintf(os.Stderr, "Can't find user: %v\n", err)
	}
	if tools.CheckPasswordHash(account.Password, password_hash) {
		return id, nil
	}
	return "", nil
}
