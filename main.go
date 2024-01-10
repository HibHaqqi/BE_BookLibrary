package main

import (
	"book-library/models"
	"book-library/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	models.ConnectDatabase()
	route := gin.Default()
	routes.SetupRoutes(route)

	port := ":8000"
	route.Run(port)
}
