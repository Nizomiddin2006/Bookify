package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Book struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}


type Comment struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	BookID  uint   `json:"book_id"`
	Content string `json:"content"`
}

type Service struct {
	DB *gorm.DB
}


func NewService(db *gorm.DB) *Service {
	
	db.AutoMigrate(&Book{}, &User{}, &Comment{})
	return &Service{DB: db}
}

func (s *Service) RegisterRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		admin.POST("/books", s.AddBook)
		admin.PUT("/books/:id", s.EditBook)
		admin.DELETE("/books/:id", s.DeleteBook)
	}

	user := router.Group("/")
	{
		user.POST("/register", s.RegisterUser)
		user.POST("/login", s.LoginUser)
		user.GET("/books", s.GetAllBooks)
		user.POST("/books/:id/comments", s.AddComment)
	}
}

func (s *Service) AddBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (s *Service) EditBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := s.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}


func (s *Service) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := s.DB.Delete(&Book{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}


func (s *Service) RegisterUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *Service) LoginUser(c *gin.Context) {
	var user User
	var input User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.DB.Where("username = ? AND password = ?", input.Username, input.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}


func (s *Service) GetAllBooks(c *gin.Context) {
	var books []Book
	if err := s.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (s *Service) AddComment(c *gin.Context) {
	bookID := c.Param("id")
	var comment Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookIDUint, err := strconv.ParseUint(bookID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	comment.BookID = uint(bookIDUint)
	if err := s.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}
	c.JSON(http.StatusOK, comment)
}
