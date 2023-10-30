package utils

import (
	"time"

	"github.com/arvinpaundra/go-eduworld/internal/configs"
	"github.com/golang-jwt/jwt"
)

type JWTCustomClaims struct {
	ID     string `json:"id"`
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type JWTCustomOption struct {
	ID        string
	Role      string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func GenerateToken(option *JWTCustomOption) (string, error) {
	claims := JWTCustomClaims{
		ID:   option.ID,
		Role: option.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  option.IssuedAt.Unix(),
			ExpiresAt: option.ExpiredAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.GetConfig("JWT_SECRET")))
}
