package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var key = os.Getenv("JWT_SECRET")

const tokenLifespan = 15 * 60                 // 15 minutes
const refreshTokenLifeSpan = 3 * 24 * 60 * 60 // 3 days
const issuer = "KTS"

func GenerateJWT(userId *uuid.UUID) (string, string, error) {
	now := time.Now()

	claims := &jwt.MapClaims{
		"exp": now.Add(time.Duration(tokenLifespan) * time.Second).Unix(),
		"iat": now.Unix(),
		"iss": issuer,
		"sub": userId.String(),
	}

	refreshClaims := &jwt.MapClaims{
		"exp": now.Add(time.Duration(refreshTokenLifeSpan) * time.Second).Unix(),
		"iat": now.Unix(),
		"iss": issuer,
		"sub": userId.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", "", err
	}

	signedRefreshToken, err := refreshToken.SignedString([]byte(key))
	if err != nil {
		return "", "", err
	}

	return signedToken, signedRefreshToken, err
}