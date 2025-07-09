package handlers

import (
	"net/http"
	"testing"

	"github.com/JWindy92/PerfectPint/internal/utils"
)

func TestCreateUser(t *testing.T) {
	router := SetupTestRouter()

	payload := map[string]string{
		"name":  "Test User",
		"email": "test@example.com",
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
	router := SetupTestRouter()

	params := map[string]string{
		"email": "test@example.com",
	}

	req, w := utils.MakeGETRequestWithQuery("/users", params)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d: %s", w.Code, w.Body.String())
	}
}
