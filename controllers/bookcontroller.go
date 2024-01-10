package bookcontroller

import (
	"book-library/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"Book": books})
}
func Create(c *gin.Context) {
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:       input.Title,
		Author:      input.Author,
		PublishYear: input.PublishYear,
		Description: input.Description,
	}
	models.DB.Create(&book)

	c.JSON(http.StatusCreated, gin.H{"data": book})
}
