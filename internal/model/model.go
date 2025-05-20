package model

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:user"` 
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Book struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"not null"`
	Author      string    `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	BookID    uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}