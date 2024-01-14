package service

import (
	"book-library/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// Fetch all books from the database
	var books []models.Book
	models.DB.Find(&books)

	// Create a map to store average ratings for each book
	averageRatings := make(map[uint]float64)
	totalRatings := make(map[uint]int)
	// Calculate average rating for each book
	for _, book := range books {
		averageRating, totalRating := CalculateUpdatedRatings(book.ID)
		averageRatings[book.ID] = averageRating
		totalRatings[book.ID] = totalRating
	}

	// Create a response structure to include both book details and average ratings
	var response []gin.H
	for _, book := range books {
		response = append(response, gin.H{
			"book":          book,
			"averageRating": averageRatings[book.ID],
			"totalRating":   totalRatings[book.ID],
		})
	}

	c.JSON(http.StatusOK, gin.H{"books": response})
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

// Update handles the PUT request for updating a book
func Update(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func Delete(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": "Book deleted"})
}
