package main

// @title           Go E-Commerce REST API
// @version         1.0
// @description     Core services for the e-commerce platform. Handles user authentication (JWT), dynamic product inventories, order lifecycles, and secure transaction tracking.
// @termsOfService  http://swagger.io/terms/
// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT
// @host            localhost:8080
// @BasePath        /api/v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/akhmadzaki/api-ecommerce-basic/config"
	"github.com/akhmadzaki/api-ecommerce-basic/database"
	_ "github.com/akhmadzaki/api-ecommerce-basic/docs"
	"github.com/akhmadzaki/api-ecommerce-basic/handler"
	"github.com/akhmadzaki/api-ecommerce-basic/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	database.RunMigrations(cfg)

	r := gin.Default()

	r.Use(middleware.TimeoutMiddleware(30 * time.Second))

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint
	r.GET("/api/v1/health", handler.GetApiHealth)

	r.POST("/api/v1/auth/login", handler.LoginHandler(cfg))
	r.POST("/api/v1/auth/register", handler.RegisterHandler)

	user := r.Group("/user")
	user.Use(middleware.JwtAuthMiddleware([]byte(cfg.JWTAccessSecret)))
	{
		user.GET("/me", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "get user profile"})
		})
	}

	r.Run(":" + cfg.ServerPort)
}
