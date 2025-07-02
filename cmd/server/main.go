package main

import (
	"go-auth-app/config"
	"go-auth-app/routes"
	"log"
	"os"
	"time"

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
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Allow all origins for local testing
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache preflight response for 12 hours
	}))

	// Register routes
	r.GET("/a", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Go Auth App!"})
	})
	r.Static("/uploads", "./uploads")

	routes.SetupRoutes(r)
	// Start the server
	r.Run(":" + os.Getenv("PORT"))
}
