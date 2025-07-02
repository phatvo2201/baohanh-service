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
	api.Use(middleware.AuthMiddleware())

	// User CRUD - admin only
	user := api.Group("/users")
	user.Use(middleware.AdminOnly())
	{
		user.GET("/", controllers.GetAllUsers)
		user.POST("/", controllers.Register)
		user.GET("/:id", controllers.GetUserByID)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	// ProductWarranty CRUD - admin only
	productWarranty := api.Group("/productwarranty")
	productWarranty.Use(middleware.AdminOnly())
	{
		productWarranty.POST("/", controllers.CreateProductWarranty)
		productWarranty.GET("/", controllers.GetAllProductWarranties)
		productWarranty.GET("/:id", controllers.GetProductWarrantyByID)
		productWarranty.PUT("/:id", controllers.UpdateProductWarranty)
		productWarranty.DELETE("/:id", controllers.DeleteProductWarranty)
		productWarranty.GET("/:id/repairs", controllers.ListRepairDetailsByProductWarrantyID)
	}

	// RepairWarranty CRUD - admin only
	repairWarranty := api.Group("/repairwarranty")
	repairWarranty.Use(middleware.AdminOnly())
	{
		repairWarranty.POST("/", controllers.CreateRepairWarranty)
		repairWarranty.GET("/", controllers.GetAllRepairWarranties)
		repairWarranty.GET("/:id", controllers.GetRepairWarrantyByID)
		repairWarranty.PUT("/:id", controllers.UpdateRepairWarranty)
		repairWarranty.DELETE("/:id", controllers.DeleteRepairWarranty)
		repairWarranty.GET("/:id/repairs", controllers.ListRepairDetailsByRepairWarrantyID)
	}
}
