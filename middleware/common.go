package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with a timeout based on user request context
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Update request with the timeout context
		c.Request = c.Request.WithContext(ctx)

		// Channel to signal when the handler finishes processing
		finished := make(chan struct{})

		go func() {
			defer close(finished)
			c.Next() // Execute the actual endpoint handler
		}()

		select {
		case <-finished:
			// Request completed successfully within the timeout limit
			return
		case <-ctx.Done():
			// Timeout was reached before the handler finished
			if ctx.Err() == context.DeadlineExceeded {
				c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{
					"error": "gateway timeout: the request took too long to process",
				})
			}
			return
		}
	}
}
