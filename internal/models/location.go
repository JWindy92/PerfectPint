package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	Name    string
	Address string
	Reviews []Review
}
