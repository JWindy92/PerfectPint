package handlers

import (
	"net/http"
	"testing"

	"github.com/JWindy92/PerfectPint/internal/database"
	"github.com/JWindy92/PerfectPint/internal/routes"
	"github.com/JWindy92/PerfectPint/internal/utils"
	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	db := database.SQLiteImpl{}

	// uHandler := NewDefaultUserHandler(&db)
	r := gin.Default()
	routes.RegisterRoutes(
		r,
		NewAuthPassthroughHandler(&db),
		NewDefaultUserHandler(&db),
	)
	return r
}

type MockUserService struct{}

func TestCreateUser(t *testing.T) {
	router := setupTestRouter()

	payload := map[string]string{
		"name":  "Test User2",
		"email": "test2@example.com",
	}

	req, w, err := utils.MakeJSONRequest("POST", "/users", payload)
	if err != nil {
		t.Fatalf("Failed to create JSON request: %v", err)
	}

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected 201, got %d: %s", w.Code, w.Body.String())
	}
}

func TestGetUserByEmail(t *testing.T) {
	router := setupTestRouter()

	params := map[string]string{
		"email": "test@example.com",
	}

	req, w := utils.MakeGETRequestWithQuery("/users", params)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

// func TestGetUserByEmail(t *testing.T) {
// 	router := setupTestRouter()

// 	url := fmt.Sprintf("/users?email=%s", "another@example.com")
// 	req := httptest.NewRequest(http.MethodGet, url, nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	if w.Code != http.StatusOK {
// 		t.Fatalf("Expected status 200, got %d: %v", w.Code, w.Body)
// 	}

// 	var fetchedUser models.User
// 	if err := json.Unmarshal(w.Body.Bytes(), &fetchedUser); err != nil {
// 		t.Fatal("Failed to unmarshal response")
// 	}

// 	fmt.Println(fetchedUser)

// 	// if fetchedUser.ID != user.ID {
// 	// 	t.Errorf("Expected user ID %d, got %d", user.ID, fetchedUser.ID)
// 	// }
// 	// if fetchedUser.Name != user.Name {
// 	// 	t.Errorf("Expected name %q, got %q", user.Name, fetchedUser.Name)
// 	// }
// }
