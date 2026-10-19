package handler

import (
	"net/http"
	"time"

	"github.com/akhmadzaki/api-ecommerce-basic/database"
	"github.com/gin-gonic/gin"
)

// HealthResponse represents the health check status payload
type HealthResponse struct {
	APIStatus string `json:"api_status" example:"healthy"`
	DBStatus  string `json:"db_status" example:"connected"`
	Timestamp string `json:"timestamp" example:"2026-10-19T15:12:00Z"`
}

// GetApiHealth godoc
// @Summary      Check API health
// @Description  Responds with API health status
// @Produce      json
// @Success      200  {object}  HealthResponse
// @Router       /health [get]
func GetApiHealth(c *gin.Context) {
	var dbStatus string
	dbStatus = "connected"
	err := database.DBConn.QueryRowContext(c.Request.Context(), "SELECT 'Database is healthy!'").Scan(&dbStatus)
	if err != nil {
		dbStatus = "not connected"
	}

	c.JSON(http.StatusOK, HealthResponse{
		APIStatus: "healthy",
		DBStatus:  dbStatus,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
