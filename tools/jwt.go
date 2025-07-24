package tools

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/zidnirifan/varnion-utils/authentication"
)

type Claims struct {
	UserID    uuid.UUID `json:"user_id"`
	Role      string    `json:"role"`
	LoginTime int64     `json:"login_time"`
	jwt.RegisteredClaims
}

func ValidateAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		secret := os.Getenv(authentication.ACCESS_SECRET_KEY)
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
