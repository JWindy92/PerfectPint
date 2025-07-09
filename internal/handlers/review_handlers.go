package handlers

import (
	"net/http"

	"github.com/JWindy92/PerfectPint/internal/common"
	"github.com/JWindy92/PerfectPint/internal/models"
	"github.com/JWindy92/PerfectPint/internal/services"
	"github.com/gin-gonic/gin"
)

type ReviewServiceInterface interface {
	CreateReview(review *models.Review) error
	GetReviewByID(id string) (*models.Review, error)
}

type ReviewHandler struct {
	Service ReviewServiceInterface
}

func NewDefaultReviewHandler(db common.DBInterface) *ReviewHandler {
	return &ReviewHandler{
		Service: services.NewReviewService(db.ConnectDB()),
	}
}

func (h *ReviewHandler) GetReviewByID(c *gin.Context) {
	id := c.Param("id")

	// if err := database.DB.Preload("Reviews").First(&user, id).Error; err != nil {
	review, err := h.Service.GetReviewByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	c.JSON(http.StatusOK, review)
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create review"})
		return
	}

	c.JSON(http.StatusCreated, review)
}
