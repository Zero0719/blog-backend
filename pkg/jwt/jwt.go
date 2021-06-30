package jwt

import (
	"blog-backend/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(id int) (string, error) {
	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(config.Conf.Expired),
			Issuer:    "goblog",
		},
	}

	tokenCliams := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenCliams.SignedString([]byte(config.Conf.Jwt))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenCliams, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.Jwt), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenCliams.Claims.(*Claims); ok && tokenCliams.Valid {
		return claims, nil
	}

	return nil, err
}
