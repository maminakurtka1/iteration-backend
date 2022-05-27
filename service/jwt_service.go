package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	AccountUUID string `json:"accoount_uuid"`
}

var jwtKey = []byte("my_secret_key")

func GenerateToken(account_uuid string) (string, error) {
	claims := &tokenClaims{
		AccountUUID: account_uuid,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func ParseToken(tokenString string) (*tokenClaims, error) {
	claims := &tokenClaims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, err
		}
		return claims, err
	}
	if !tkn.Valid {
		return claims, err
	}
	return claims, err
}

func RefreshToken(tokenString string) (string, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	if time.Until(time.Unix(claims.ExpiresAt, 0)) >= 30*time.Second {
		return "claims", err
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return tokenString, err
}
