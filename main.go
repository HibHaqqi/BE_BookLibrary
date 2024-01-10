package main

import (
	bookcontroller "book-library/controllers"
	"book-library/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	models.ConnectDatabase()
	route := gin.Default()
	route.GET("/api/books", bookcontroller.Index)

	port := ":8000"
	route.Run(port)
}
