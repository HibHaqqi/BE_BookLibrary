package api

import (
	authcontroller "book-library/controllers/auth"
	servicecontroller "book-library/controllers/service"

	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(route *gin.RouterGroup) {
	api := route.Group("")

	// Use the service controller
	bookRoutes := api.Group("/books")
	{
		bookRoutes.GET("", servicecontroller.Index)
		bookRoutes.POST("", servicecontroller.Create)
		bookRoutes.PUT("/:id", servicecontroller.Update)
		bookRoutes.DELETE("/:id", servicecontroller.Delete)
	}

	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/register", authcontroller.Register)
		authRoutes.POST("/login", authcontroller.Login)
	}

	// Add more API routes as needed
}
