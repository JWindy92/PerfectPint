package services

import (
	"testing"

	"github.com/JWindy92/PerfectPint/internal/models"
	"github.com/JWindy92/PerfectPint/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetUserByEmail(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("../../perfectpints.db?_cache=shared"), &gorm.Config{})
	db.AutoMigrate(&models.User{})

	service := NewUserService(db)
	testUser := models.User{Name: "Another User", Email: "another@example.com"}

	user, err := service.GetUserByEmail("another@example.com")
	if err != nil {
		t.Fatalf("User not found: %v", err)
	}

	if user.Name != testUser.Name {
		t.Errorf("Expected name %s, got %s", testUser.Name, user.Name)
	}

	if user.Email != testUser.Email {
		t.Errorf("Expected email %s, got %s", testUser.Email, user.Email)
	}

	utils.PrettyPrint(user)
}
