package services

import (
	"log"

	"github.com/JWindy92/PerfectPint/internal/models"
	"github.com/JWindy92/PerfectPint/internal/utils"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(conn *gorm.DB) *UserService {
	return &UserService{DB: conn}
}

func (s *UserService) CreateUser(user *models.User) error {
	utils.PrettyPrint(user)
	return s.DB.Create(user).Error
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	log.Printf("Attempting to find user with email: %s", email)
	var user models.User
	if err := s.DB.Preload("Reviews").Where("email = ?", email).First(&user).Error; err != nil {
		log.Println("No user found")
		return nil, err
	}
	log.Println("User found")
	return &user, nil
}
