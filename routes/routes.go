package routes

import (
	bookcontroller "book-library/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine) {
	api := route.Group("/api")

	// Use the book controller
	bookRoutes := api.Group("/books")
	{
		bookRoutes.GET("", bookcontroller.Index)
		bookRoutes.POST("", bookcontroller.Create)
	}
}
