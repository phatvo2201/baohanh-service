package routes

import (
	"go-auth-app/controllers"
	"go-auth-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.GET("/me", middleware.AuthMiddleware(), controllers.GetCurrentUser)
	}

	api := router.Group("/api")

	// User CRUD - admin only
	user := api.Group("/users")
	{
		user.GET("/", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.GetAllUsers)
		user.POST("/", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.Register)
		user.GET("/:id", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.GetUserByID)
		user.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.UpdateUser)
		user.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.DeleteUser)
	}

	// ProductWarranty
	productWarranty := api.Group("/productwarranty")
	{
		// Public endpoints (no token required)
		productWarranty.GET("/", controllers.GetAllProductWarranties)
		productWarranty.GET("/:id", controllers.GetProductWarrantyByID)
		productWarranty.GET("/:id/repairs", controllers.ListRepairDetailsByProductWarrantyID)

		// Protected endpoints (admin only)
		productWarranty.POST("/", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.CreateProductWarranty)
		productWarranty.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.UpdateProductWarranty)
		productWarranty.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.DeleteProductWarranty)
	}

	// RepairWarranty
	repairWarranty := api.Group("/repairwarranty")
	{
		// Public endpoints (no token required)
		repairWarranty.GET("/", controllers.GetAllRepairWarranties)
		repairWarranty.GET("/:id", controllers.GetRepairWarrantyByID)
		repairWarranty.GET("/:id/repairs", controllers.ListRepairDetailsByRepairWarrantyID)

		// Protected endpoints (admin only)
		repairWarranty.POST("/", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.CreateRepairWarranty)
		repairWarranty.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.UpdateRepairWarranty)
		repairWarranty.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminOnly(), controllers.DeleteRepairWarranty)
	}
}
