package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/akhmadzaki/api-ecommerce-basic/config"
	"github.com/akhmadzaki/api-ecommerce-basic/database"
	"github.com/akhmadzaki/api-ecommerce-basic/model"
	"github.com/akhmadzaki/api-ecommerce-basic/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// RegisterHandler godoc
// @Summary      Register user
// @Description  Register new user with email, password, first name, and last name
// @Accept      json
// @Produce      json
// @Param        request  body  model.RegisterRequest  true  "Register Request"
// @Success      200  {object}  model.RegisterResponse
// @Router       /auth/register [post]
func RegisterHandler(c *gin.Context) {
	var registerReq model.RegisterRequest

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	passwordHash, err := utils.HashPassword(registerReq.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var userId int64
	query := `INSERT INTO users(email, password_hash, first_name, last_name) VALUES($1, $2, $3, $4) RETURNING id`
	err = database.DBConn.QueryRowContext(c.Request.Context(), query, registerReq.Email, passwordHash, registerReq.FirstName, registerReq.LastName).Scan(&userId)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
				return
			}
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, model.RegisterResponse{
		Message: "User created",
		UserId:  userId,
	})
}

// LoginHandler godoc
// @Summary     User login
// @Description  User login with email and password
// @Accept      json
// @Produce      json
// @Param        request  body  model.LoginRequest  true  "Login credentials"
// @Success      200  {object}  model.LoginResponse
// @Router       /auth/login [post]
func LoginHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginReq model.LoginRequest

		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		query := `SELECT id, role, password_hash FROM users WHERE email = $1`
		var userId int64
		var role string
		var passwordFromDB string
		err := database.DBConn.QueryRowContext(c.Request.Context(), query, loginReq.Email).Scan(&userId, &role, &passwordFromDB)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		is_verified, err := utils.VerifyPassword(loginReq.Password, passwordFromDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		if !is_verified {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
			return
		}

		uuid := uuid.New().String()

		accessToken, err := utils.GenerateAccessToken(userId, role, []byte(cfg.JWTAccessSecret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		refreshToken, refreshTokenHash, expiredAt, err := utils.GenerateRefreshToken(userId, uuid, []byte(cfg.JWTRefreshSecret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		query = `INSERT INTO refresh_tokens(user_id, token_hash, is_revoked, expires_at) VALUES($1, $2, $3, $4) RETURNING id`
		var tokenId int64
		err = database.DBConn.QueryRowContext(c.Request.Context(), query, userId, refreshTokenHash, false, expiredAt).Scan(&tokenId)
		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) {
				if pqErr.Code == "23505" {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Token already exists"})
					return
				}
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, model.LoginResponse{
			Message:      "Login successful",
			UserId:       userId,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
	}

}
