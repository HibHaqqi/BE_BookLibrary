package service

import (
	"book-library/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SubmitRating handles the submission of user ratings.
func SubmitRating(c *gin.Context) {
	var input models.Rating
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save rating to the database
	createRating(input)

	// Calculate updated average rating and total ratings
	averageRating, totalRatings := calculateUpdatedRatings(input.BookId)

	c.JSON(http.StatusOK, gin.H{
		"averageRating": averageRating,
		"totalRatings":  totalRatings,
	})
}

// CreateRating creates a new rating in the database.
func createRating(input models.Rating) {
	newRating := models.Rating{
		MemberId: input.MemberId,
		BookId:   input.BookId,
		Rating:   input.Rating,
	}
	models.DB.Create(&newRating)
}

// CalculateUpdatedRatings calculates the updated average rating and total ratings for a book.
func calculateUpdatedRatings(bookId uint) (float64, int) {
	var ratings []models.Rating
	result := models.DB.Where("book_id = ?", bookId).Find(&ratings)
	if result.Error != nil {
		// Handle error
		return 0, 0
	}

	var totalRatings int
	var sumRatings float64

	for _, r := range ratings {
		sumRatings += r.Rating
		totalRatings++
	}

	if totalRatings > 0 {
		averageRating := sumRatings / float64(totalRatings)
		return averageRating, totalRatings
	}

	return 0, 0
}
