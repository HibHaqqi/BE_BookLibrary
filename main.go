package main

import (
	"book-library/models"
	"book-library/routes/api"
	"book-library/routes/view"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	models.ConnectDatabase()
	route := gin.Default()

	route.LoadHTMLGlob("templates/*.html")
	// Serve static files
	route.Static("/static", "./static")
	// Set up API routes
	api.SetupAPIRoutes(route.Group("/api"))

	// Set up view routes
	view.SetupViewRoutes(route.Group(""))

	port := ":8000"
	route.Run(port)
}
