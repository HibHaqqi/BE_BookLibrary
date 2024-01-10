package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupViewRoutes(route *gin.RouterGroup) {
	views := route.Group("")

	// Add your view routes as needed
	views.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.html", gin.H{"Title": "Home Page"})
	})

}
