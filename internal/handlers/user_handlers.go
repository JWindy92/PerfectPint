package handlers

import (
	"net/http"

	"github.com/JWindy92/PerfectPint/internal/common"
	"github.com/JWindy92/PerfectPint/internal/models"
	"github.com/JWindy92/PerfectPint/internal/services"
	"github.com/JWindy92/PerfectPint/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	GetUserByID(string) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	CreateUser(*models.User) error
}
type UserHandler struct {
	Service UserServiceInterface
}

func NewDefaultUserHandler(db common.DBInterface) *UserHandler {
	return &UserHandler{
		Service: services.NewUserService(db.ConnectDB()),
	}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	// if err := database.DB.Preload("Reviews").First(&user, id).Error; err != nil {
	user, err := h.Service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// TODO: Check for SQL injection (does gorm handle it natively?)
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Query("email") // or use Param() if you're using path params

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	user, err := h.Service.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.PrettyPrint(req)

	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hash(req.Password), // replace with actual hash function
	}

	if err := h.Service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// func hash(pw string) string {
// 	return pw
// }
