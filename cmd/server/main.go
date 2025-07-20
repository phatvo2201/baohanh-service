package main

import (
	"go-auth-app/config"
	"go-auth-app/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	// Configure CORS to allow all origins for local development
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept", "X-Requested-With", "X-CSRF-Token"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = false // Set to false when using AllowAllOrigins
	config.MaxAge = 12 * 60 * 60    // 12 hours

	r.Use(cors.New(config))

	// Register routes
	r.GET("/a", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Go Auth App!"})
	})
	r.Static("/uploads", "./uploads")

	routes.SetupRoutes(r)

	// Start the server
	r.Run(":" + os.Getenv("PORT"))
}
