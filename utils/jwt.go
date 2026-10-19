package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccessClaims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserID  int64  `json:"user_id"`
	TokenID string `json:"token_id"` // Matches the unique identifier in your database table
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID int64, role string, jwtAccessSecret []byte) (string, error) {
	claims := AccessClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtAccessSecret)
}

func GenerateRefreshToken(userID int64, tokenID string, jwtRefreshSecret []byte) (string, string, time.Time, error) {
	uuid := uuid.New().String()
	currentTime := time.Now()
	claims := RefreshClaims{
		UserID:  userID,
		TokenID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(currentTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtRefreshSecret)
	if err != nil {
		return "", "", time.Time{}, err
	}
	tokenBytes := sha256.Sum256([]byte(tokenString))
	tokenHash := hex.EncodeToString(tokenBytes[:])

	return tokenString, tokenHash, currentTime.Add(7 * 24 * time.Hour), nil
}

func ValidateRefreshToken(tokenStr string, jwtRefreshSecret []byte) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtRefreshSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	claims, ok := token.Claims.(*RefreshClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims, nil
}
