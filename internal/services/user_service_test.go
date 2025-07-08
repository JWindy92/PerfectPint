package services

import (
	"testing"

	"github.com/JWindy92/PerfectPint/internal/database"
	"github.com/JWindy92/PerfectPint/internal/models"
	"github.com/JWindy92/PerfectPint/internal/utils"
)

func TestGetUserByEmail(t *testing.T) {
	db := database.SQLiteImpl{}
	conn := db.ConnectDB()
	service := NewUserService(conn)
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

func TestCreateUser(t *testing.T) {
	db := database.SQLiteImpl{}
	conn := db.ConnectDB()

	service := NewUserService(conn)
	testUser := models.User{Name: "One More User", Email: "onemore@example.com"}
	// testUser := models.User{Name: "Another User", Email: "another@example.com"}

	if err := service.CreateUser(&testUser); err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
}
