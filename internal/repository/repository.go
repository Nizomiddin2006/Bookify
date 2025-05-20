package repository

import (
	"gorm.io/gorm"
	"Bookify/internal/model"
)

type BookRepository interface {
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book) error
	DeleteBook(id uint) error
	GetAllBooks() ([]model.Book, error)
	GetBookByID(id uint) (*model.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) CreateBook(book *model.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) UpdateBook(book *model.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) DeleteBook(id uint) error {
	return r.db.Delete(&model.Book{}, id).Error
}

func (r *bookRepository) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) GetBookByID(id uint) (*model.Book, error) {
	var book model.Book
	err := r.db.First(&book, id).Error
	return &book, err
}
