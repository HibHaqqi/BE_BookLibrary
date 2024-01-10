package view

import (
	"github.com/gin-gonic/gin"
)

func SetupViewRoutes(route *gin.RouterGroup) {
	views := route.Group("")

	// Add your view routes as needed
	views.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the home page"})
	})

}
