package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtUtils struct {
	Secret   string
	TokenTTL time.Duration
}

type UserClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	UserRole string `json:"user_role"`
}

type tokenClaims[T any] struct {
	Claims T
	jwt.RegisteredClaims
}

func (j *JwtUtils) GenerateToken(userClaims UserClaims) (string, error) {
	now := time.Now()
	claims := tokenClaims[UserClaims]{
		Claims: userClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(j.TokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tok.SignedString([]byte(j.Secret))
}

func (j *JwtUtils) VerifyToken(token string) (UserClaims, error) {
	token = strings.TrimSpace(token)
	if token == "" {
		return UserClaims{}, errors.New("token is empty")
	}

	var claims tokenClaims[UserClaims]
	parsed, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return UserClaims{}, err
	}
	if parsed == nil || !parsed.Valid {
		return UserClaims{}, errors.New("invalid token")
	}
	return claims.Claims, nil
}
