package main

import (
	"go-auth-app/config"
	"go-auth-app/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Custom CORS middleware that allows everything for development
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Allow all origins for testing
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
		}

		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Length, Content-Type, Authorization, Accept, X-Requested-With, X-CSRF-Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
		c.Header("Access-Control-Max-Age", "43200") // 12 hours

		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to the database
	config.ConnectDatabase()
	// config.SeedDatabase()
	config.InitGoogleOAuth()

	// Initialize Gin router
	r := gin.Default()

	// Use our custom permissive CORS middleware (works with all origins including null/file://)
	r.Use(CORSMiddleware())

	// Register routes
	r.GET("/a", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Go Auth App!"})
	})
	r.Static("/uploads", "./uploads")

	routes.SetupRoutes(r)

	// Start the server
	r.Run(":" + os.Getenv("PORT"))
}
