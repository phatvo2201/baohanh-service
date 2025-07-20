package main

import (
	"go-auth-app/config"
	"go-auth-app/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow all origins
		c.Header("Access-Control-Allow-Origin", "*")

		c.Header("Access-Control-Allow-Headers", "*")

		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// Continue to the next middleware/handler
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

	// Use our custom CORS middleware
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
