package tools

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/varnion-rnd/utils/authentication"
	"github.com/varnion-rnd/utils/permission"
)

type Claims struct {
	UserID    uuid.UUID `json:"user_id"`
	Role      string    `json:"role"`
	LoginTime int64     `json:"login_time"`
	jwt.RegisteredClaims
}

type PermissionClaims struct {
	Data []*PermissionToken `json:"data"`
	jwt.RegisteredClaims
}

type PermissionToken struct {
	AppName string  `json:"app_name"`
	Menus   []*Menu `json:"menus"`
}

type Menu struct {
	MenuName   string    `json:"menu_name"`
	Permission []*string `json:"permission"`
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

func ValidatePermissionToken(tokenString string) (*PermissionClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &PermissionClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := os.Getenv(permission.PERMISSION_SECRET_KEY)
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*PermissionClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
