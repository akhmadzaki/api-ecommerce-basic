package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/akhmadzaki/api-ecommerce-basic/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthMiddleware(jwtAccessSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ""
		bearerToken := c.GetHeader("Authorization")

		if len(bearerToken) > 7 && strings.EqualFold(bearerToken[0:7], "BEARER ") {
			tokenString = bearerToken[7:]
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &utils.AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing algorithm: %v", token.Header["alg"])
			}
			return jwtAccessSecret, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(*utils.AccessClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		userID := claims.UserID

		c.Set("user_id", userID)
		c.Next()
	}
}
