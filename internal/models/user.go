package models

import "gorm.io/gorm"

// type UserRecord struct {
// 	User
// 	PasswordHash string
// }

// func (UserRecord) TableName() string {
// 	return "users"
// }

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"unique"`
	PasswordHash string
	// Reviews []Review
}
