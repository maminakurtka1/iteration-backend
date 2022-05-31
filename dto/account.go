package dto

import (
	"time"
)

type Account struct {
	Id         string    `db:"id"`
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type AccountSignUp struct {
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type AccountSignIn struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type AccountIdentities struct {
	Id           int32     `db:"id"`
	AccountId    string    `db:"account_id"`
	InsertedAt   time.Time `db:"inserted_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	PhoneNumber  string    `db:"phone_number"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CityId       int32     `db:"city_id"`
}
