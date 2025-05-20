package main

import (
	"Bookify/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) AddBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Book added"})
}

func (h *Handler) EditBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Book edited"})
}

func (h *Handler) DeleteBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func (h *Handler) RegisterUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func (h *Handler) LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
}

func (h *Handler) GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"books": []string{}})
}

func (h *Handler) AddComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Comment added"})
}

func main() {
	svc := &service.Service{}
	handler := NewHandler(svc)
	router := gin.Default()

	router.POST("/books", handler.AddBook)
	router.PUT("/books/:id", handler.EditBook)
	router.DELETE("/books/:id", handler.DeleteBook)
	router.POST("/users/register", handler.RegisterUser)
	router.POST("/users/login", handler.LoginUser)
	router.GET("/books", handler.GetAllBooks)
	router.POST("/comments", handler.AddComment)

	router.Run("localhost:8080")
}
