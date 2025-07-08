package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JWindy92/PerfectPint/internal/database"
	"github.com/JWindy92/PerfectPint/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	db, _ := gorm.Open(sqlite.Open("../../perfectpints.db?_cache=shared"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	database.DB = db

	uHandler := NewDefaultUserHandler()
	r := gin.Default()
	r.POST("/users", uHandler.CreateUser)
	r.GET("/users/:id", uHandler.GetUserByID)
	return r
}

func TestCreateUser(t *testing.T) {
	router := setupTestRouter()

	userPayload := map[string]string{
		"name":  "Test User",
		"email": "test@example.com",
	}
	body, _ := json.Marshal(userPayload)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201, got %d", w.Code)
	}

	var createdUser models.User
	if err := json.Unmarshal(w.Body.Bytes(), &createdUser); err != nil {
		t.Fatal("Failed to unmarshal response")
	}

	if createdUser.Name != userPayload["name"] {
		t.Errorf("Expected name %q, got %q", userPayload["name"], createdUser.Name)
	}
	if createdUser.Email != userPayload["email"] {
		t.Errorf("Expected email %q, got %q", userPayload["email"], createdUser.Email)
	}
}

func TestGetUserByID(t *testing.T) {
	router := setupTestRouter()

	// Create a user directly in DB
	// user := models.User{Name: "Another User", Email: "another@example.com"}
	// database.DB.Create(&user)
	// user := models.User{}
	url := fmt.Sprintf("/users/%d", 2)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d: %v", w.Code, w.Body)
	}

	var fetchedUser models.User
	if err := json.Unmarshal(w.Body.Bytes(), &fetchedUser); err != nil {
		t.Fatal("Failed to unmarshal response")
	}

	fmt.Println(fetchedUser)

	// if fetchedUser.ID != user.ID {
	// 	t.Errorf("Expected user ID %d, got %d", user.ID, fetchedUser.ID)
	// }
	// if fetchedUser.Name != user.Name {
	// 	t.Errorf("Expected name %q, got %q", user.Name, fetchedUser.Name)
	// }
}
