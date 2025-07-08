package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/JWindy92/PerfectPint/internal/models"
	"github.com/JWindy92/PerfectPint/internal/utils"
)

func TestLogin(t *testing.T) {
	router := SetupTestRouter()

	payload := map[string]string{
		"email":    "auth@example.com",
		"password": "12345",
	}

	req, w, err := utils.MakeJSONRequest("POST", "/login", payload)
	if err != nil {
		t.Fatalf("Failed to create JSON request: %v", err)
	}

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d: %s", w.Code, w.Body.String())
	}
	log.Println("Req body: ", w.Body.String())
	var resp models.LoginSuccess
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatal("Failed to parse JSON response")
	}

	utils.PrettyPrint(resp)
}
