package services

import (
	"strconv"
	"testing"

	"github.com/JWindy92/PerfectPint/internal/database"
	"github.com/JWindy92/PerfectPint/internal/models"
	"github.com/JWindy92/PerfectPint/internal/utils"
)

func TestGetLocationByID(t *testing.T) {
	db := database.SQLiteImpl{}
	conn := db.ConnectDB()
	service := NewLocationService(conn)
	testLoc := models.Location{
		ID:      1,
		Name:    "Cool Bar",
		Address: "123 Bar Street", // assumes user with ID 1 exists
	}

	loc, err := service.GetLocationByID(strconv.FormatUint(uint64(testLoc.ID), 10))
	if err != nil {
		t.Fatalf("Review not found: %v", err)
	}

	if loc.Name != testLoc.Name {
		t.Errorf("Expected name %s, got %s", testLoc.Name, loc.Name)
	}

	if loc.Address != testLoc.Address {
		t.Errorf("Expected email %s, got %s", testLoc.Address, loc.Address)
	}

	utils.PrettyPrint(loc)
}

func TestCreateLocation(t *testing.T) {
	db := database.SQLiteImpl{}
	conn := db.ConnectDB()

	service := NewLocationService(conn)
	loc := models.Location{
		Name:    "Cool Bar",
		Address: "123 Bar Street", // assumes user with ID 1 exists
	}

	if err := service.CreateLocation(&loc); err != nil {
		t.Fatalf("Failed to create review: %v", err)
	}
}
