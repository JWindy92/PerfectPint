package services

import (
	"strconv"
	"testing"

	"github.com/JWindy92/PerfectPint/internal/database"
	"github.com/JWindy92/PerfectPint/internal/models"
	"github.com/JWindy92/PerfectPint/internal/utils"
)

func TestGetReviewByID(t *testing.T) {
	db := database.SQLiteImpl{}
	conn := db.ConnectDB()
	service := NewReviewService(conn)
	testReview := models.Review{
		ID:         1,
		Rating:     5,
		Comment:    "Great Beer!",
		UserID:     1,
		LocationID: 1,
	}

	review, err := service.GetReviewByID(strconv.FormatUint(uint64(testReview.ID), 10))
	if err != nil {
		t.Fatalf("Review not found: %v", err)
	}

	if review.UserID != testReview.UserID {
		t.Errorf("Expected name %d, got %d", testReview.UserID, review.UserID)
	}

	if review.LocationID != testReview.LocationID {
		t.Errorf("Expected email %d, got %d", testReview.LocationID, review.LocationID)
	}

	utils.PrettyPrint(review)
}

func TestCreateReview(t *testing.T) {
	db := database.SQLiteImpl{}
	conn := db.ConnectDB()

	service := NewReviewService(conn)
	review := models.Review{
		Rating:     4,
		Comment:    "Great beer!",
		UserID:     1, // assumes user with ID 1 exists
		LocationID: 1, // assumes location with ID 1 exists
	}

	if err := service.CreateReview(&review); err != nil {
		t.Fatalf("Failed to create review: %v", err)
	}
}
