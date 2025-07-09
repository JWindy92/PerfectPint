package services

import (
	"github.com/JWindy92/PerfectPint/internal/models"
	"gorm.io/gorm"
)

type LocationService struct {
	DB *gorm.DB
}

func NewLocationService(conn *gorm.DB) *LocationService {
	return &LocationService{DB: conn}
}

func (s *LocationService) CreateLocation(loc *models.Location) error {
	return s.DB.Create(loc).Error
}

func (s *LocationService) GetLocationByID(id string) (*models.Location, error) {
	var loc models.Location
	if err := s.DB.Preload("Reviews").First(&loc, id).Error; err != nil {
		return nil, err
	}
	return &loc, nil
}
