// pkg/utils/middleware.go
package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger is a middleware that logs HTTP requests
func RequestLogger(logger Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log request
		timestamp := time.Now()
		latency := timestamp.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		logger.Info("Request",
			"timestamp", timestamp.Format(time.RFC3339),
			"method", method,
			"path", path,
			"status", statusCode,
			"latency", latency,
			"ip", clientIP,
			"error", errorMessage,
		)
	}
}

// CORSMiddleware handles Cross-Origin Resource Sharing (CORS)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// RateLimiterMiddleware is a basic rate limiter middleware
// Note: For production, consider using a more sophisticated rate limiter
func RateLimiterMiddleware(logger Logger, rps int, burst int) gin.HandlerFunc {
	// Simple in-memory store for IP-based rate limiting
	// For production, consider using Redis or another distributed store
	limiters := make(map[string]time.Time)
	
	return func(c *gin.Context) {
		ip := c.ClientIP()
		
		// Check if IP is in the limiter map
		lastReq, exists := limiters[ip]
		now := time.Now()
		
		// If IP exists and request is too soon
		if exists && now.Sub(lastReq) < time.Second/time.Duration(rps) {
			logger.Warn("Rate limit exceeded",
				"ip", ip,
				"path", c.Request.URL.Path,
			)
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Please try again later.",
			})
			c.Abort()
			return
		}
		
		// Update last request time
		limiters[ip] = now
		
		// Clean up old entries periodically (simple garbage collection)
		// For production, implement a proper cleanup mechanism
		if len(limiters) > 10000 {
			for k, v := range limiters {
				if now.Sub(v) > time.Minute {
					delete(limiters, k)
				}
			}
		}
		
		c.Next()
	}
}

// AuthMiddleware is a placeholder for authentication middleware
// Replace with your actual authentication logic
func AuthMiddleware(logger Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		
		// For this example, just check if token starts with "Bearer "
		// In a real application, validate the token properly
		if len(token) < 7 || token[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}
		
		// Token validation would go here
		// ...
		
		// Set user identifier for the request
		c.Set("user_id", "sample_user_id")
		
		c.Next()
	}
}