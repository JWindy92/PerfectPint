package handlers

import (
	"net/http"

	"github.com/JWindy92/PerfectPint/internal/common"
	"github.com/JWindy92/PerfectPint/internal/models"
	"github.com/JWindy92/PerfectPint/internal/services"
	"github.com/gin-gonic/gin"
)

type LocationServiceInterface interface {
	CreateLocation(loc *models.Location) error
	GetLocationByID(id string) (*models.Location, error)
}

type LocationHandler struct {
	Service LocationServiceInterface
}

func NewDefaultLocationHandler(db common.DBInterface) *LocationHandler {
	return &LocationHandler{
		Service: services.NewLocationService(db.ConnectDB()),
	}
}

func (h *LocationHandler) GetLocationByID(c *gin.Context) {
	id := c.Param("id")

	// if err := database.DB.Preload("Reviews").First(&user, id).Error; err != nil {
	loc, err := h.Service.GetLocationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	c.JSON(http.StatusOK, loc)
}

func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var loc models.Location
	if err := c.ShouldBindJSON(&loc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateLocation(&loc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create location"})
		return
	}

	c.JSON(http.StatusCreated, loc)
}
