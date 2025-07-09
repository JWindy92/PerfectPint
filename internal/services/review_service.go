package services

import (
	"github.com/JWindy92/PerfectPint/internal/models"
	"gorm.io/gorm"
)

type ReviewService struct {
	DB *gorm.DB
}

func NewReviewService(conn *gorm.DB) *ReviewService {
	return &ReviewService{DB: conn}
}

func (s *ReviewService) CreateReview(review *models.Review) error {
	return s.DB.Create(review).Error
}

func (s *ReviewService) GetReviewByID(id string) (*models.Review, error) {
	var review models.Review
	if err := s.DB.First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}
