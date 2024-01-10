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
