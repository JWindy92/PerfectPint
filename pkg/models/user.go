package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

//TODO: Validate incoming json, i.e. role should be only "admin" or "user"
