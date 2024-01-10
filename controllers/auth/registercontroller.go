package auth

import (
	"book-library/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input models.Member

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email already exists
	if isExistingEmail(input.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Create a new member with hashed password
	member := models.Member{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: string(hashedPassword), // Store hashed password in the database
		Role:     input.Role,
	}

	// Save the member to the database
	if err := models.DB.Create(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating member"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": member})

}

// Fungsi bantuan untuk memeriksa apakah email sudah ada di database
func isExistingEmail(email string) bool {
	var existingMember models.Member
	if err := models.DB.Where("email = ?", email).First(&existingMember).Error; err != nil {
		return false // Email tidak ditemukan
	}
	return true // Email sudah ada
}
