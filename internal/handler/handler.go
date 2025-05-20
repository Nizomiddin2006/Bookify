package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		admin.POST("/books", AddBook)
		admin.PUT("/books/:id", EditBook)
		admin.DELETE("/books/:id", DeleteBook)
	}
	router.POST("/register", RegisterUser)
	router.POST("/login", LoginUser)
	router.GET("/books", GetAllBooks)
	router.POST("/books/:id/comments", AddComment)
}

func AddBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Book added"})
}

func EditBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Book edited"})
}

func DeleteBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func RegisterUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
}

func GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of books"})
}

func AddComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Comment added"})
}