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
	FirstName        string `gorm:"type:varchar(255);not null"`
	LastName         string `gorm:"type:varchar(255);not null"`
	Email            string `gorm:"uniqueIndex;not null"`
	PasswordHash     string `gorm:"not null"`
	Role             string `gorm:"type:varchar(255);not null"`
	VerificationCode string
	Verified         bool `gorm:"not null"`
}

/*
JSON Input structures to be used by the API.
*/
type UserSignUpInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UserSignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

/*
JSON Output structures to be used by the API.
*/
type UserDataResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
