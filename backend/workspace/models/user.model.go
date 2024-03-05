package models

import (
	"time"

	"gorm.io/gorm"
)

/*
Model structures to be used for communicating with the API.
*/
type User struct {
	gorm.Model
	ID           uint   `gorm:"primary_key"`
	FirstName    string `gorm:"type:varchar(255);not null"`
	LastName     string `gorm:"type:varchar(255); not null"`
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"type:varchar(255);;not null"`
	Verified     bool   `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Token        Token
}

type Token struct {
	gorm.Model
	ID        uint      `gorm:"primary_key"`
	TokenHash string    `gorm:"type:varchar(255);not null"`
	Expiry    time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	UserID    uint
}

/*
JSON structures to be used by the API.
*/
type UserSignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
