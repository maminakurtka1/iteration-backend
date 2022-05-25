package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	AccountUUID string `json:"accoount_uuid"`
}

func GenerateToken(account_uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix()},
		AccountUUID: account_uuid})
	// TODO: Take out key from func to .env or config file
	return token.SignedString([]byte("key"))
}
