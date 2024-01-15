package service

import (
	"book-library/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookInfo struct {
	ID            uint
	Title         string
	Author        string
	PublishYear   int
	Description   string
	AverageRating float64
}

func Index(c *gin.Context) {
	// Fetch all books from the database
	var books []models.Book
	models.DB.Find(&books)

	var bookInfoList []BookInfo
	for _, book := range books {
		// Capture both values returned by CalculateUpdatedRatings
		averageRating, _ := CalculateUpdatedRatings(book.ID)

		// Populate BookInfo struct with necessary fields
		bookInfo := BookInfo{
			ID:            book.ID,
			Title:         book.Title,
			Author:        book.Author,
			PublishYear:   book.PublishYear,
			Description:   book.Description,
			AverageRating: averageRating,
		}

		// Append to the slice
		bookInfoList = append(bookInfoList, bookInfo)
	}

	// Render HTML template with the book information
	c.JSON(http.StatusOK, gin.H{
		"Books": bookInfoList,
	})
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
