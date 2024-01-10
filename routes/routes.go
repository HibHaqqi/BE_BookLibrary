package routes

import (
	authcontroller "book-library/controllers/auth"
	bookcontroller "book-library/controllers/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine) {
	api := route.Group("/api")

	// Use the book controller
	bookRoutes := api.Group("/books")
	{
		bookRoutes.GET("", bookcontroller.Index)
		bookRoutes.POST("", bookcontroller.Create)
		bookRoutes.PUT("/:id", bookcontroller.Update)
		bookRoutes.DELETE("/:id", bookcontroller.Delete)
	}

	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("", authcontroller.Register)
		authRoutes.POST("/login", authcontroller.Login)

	}
}
