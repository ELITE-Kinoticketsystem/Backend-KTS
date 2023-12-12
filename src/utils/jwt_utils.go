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
const leeway = 5 * 60                         // 5 minutes
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

func ValidateToken(tokenString string) (*uuid.UUID, error) {
	validMethodsOption := jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()})
	leewayOption := jwt.WithLeeway(time.Duration(leeway) * time.Second)
	issuerOption := jwt.WithIssuer(issuer)
	issuedAtOption := jwt.WithIssuedAt()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	}, validMethodsOption, leewayOption, issuerOption, issuedAtOption)

	if err != nil {
		return nil, err
	}

	var claims jwt.MapClaims
	var ok bool

	if claims, ok = token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, jwt.ErrTokenUnverifiable
	}

	userId, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return nil, err
	}

	return &userId, nil
}

func ExtractToken(authHeader string) (string, error) {
	if len(authHeader) < 7 {
		return "", jwt.ErrInvalidKey
	}

	// Return token without "Bearer " prefix
	return authHeader[7:], nil
}
