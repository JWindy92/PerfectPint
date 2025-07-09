package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	Comment    string
	Rating     int
	UserID     uint
	LocationID uint
}
