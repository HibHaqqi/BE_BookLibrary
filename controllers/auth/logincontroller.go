package auth

import (
	"book-library/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input models.Member

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the member by email
	var member models.Member
	if err := models.DB.Where("email = ?", input.Email).First(&member).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email not register"})
		return
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// If email and password are valid, generate a token (JWT) if needed
	// You can use a JWT library to generate a token and send it as a response

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
