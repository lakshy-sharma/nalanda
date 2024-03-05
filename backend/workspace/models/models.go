package models

import "gorm.io/gorm"

var Db *gorm.DB

// This function takes a DB connection as an input and provides you with the models for communicating with the data.
func New(dbConn *gorm.DB) Models {
	Db = dbConn

	return Models{
		User:  User{},
		Token: Token{},
	}
}

// An encapsulation struct for models of an application.
type Models struct {
	User  User
	Token Token
}
