package service

import (
	"time"

	"github.com/apex/log"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	AccountId string `json:"accoount_id"`
}

var jwtKey = []byte("my_secret_key")

func GenerateToken(account_id string) (string, error) {
	claims := &tokenClaims{
		AccountId: account_id,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.WithError(err).Error("Can't create new token!")
	}
	return tokenString, err
}

func ParseToken(tokenString string) (*tokenClaims, error) {
	claims := &tokenClaims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.WithError(err).Error("Can't parse token!")
			return claims, err
		}
		log.WithError(err).Error("Can't parse token!")
		return claims, err
	}
	if !tkn.Valid {
		log.WithError(err).Error("Not valid token!")
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
		log.WithError(err).Error("Not valid token!")
		return "", err
	}
	expirationTime := time.Now().Add(10 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	if err != nil {
		log.WithError(err).Error("Can't recreate token!")
	}
	return tokenString, err
}
