package Controllers

import (
	"mvc-api/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := Models.Book{Title: input.Title, Author: input.Author}
	// models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
