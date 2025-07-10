package handlers

import (
	"fmt"
	"net/http"

	"github.com/JWindy92/PerfectPint/internal/common"
	"github.com/JWindy92/PerfectPint/internal/models"
	"github.com/JWindy92/PerfectPint/internal/services"
	"github.com/JWindy92/PerfectPint/internal/utils"
	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                    Basic Auth Passthrough Implementation                   */
/* -------------------------------------------------------------------------- */

type UserAuthInterface interface {
	GetUserByEmail(string) (*models.User, error)
	CreateUser(*models.User) error
}

// Basic implementation of the routes.AuthInterface to allow for dev work without full authentication logic
// ! Not for production
type AuthPassthroughHandler struct {
	UserService UserAuthInterface
}

func NewAuthPassthroughHandler(db common.DBInterface) *AuthPassthroughHandler {
	return &AuthPassthroughHandler{
		UserService: services.NewUserService(db.ConnectDB()),
	}
}

func (a *AuthPassthroughHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	utils.PrettyPrint(req)

	user, err := a.UserService.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("unexpected error: %v", err)})
	}

	if a.comparePassword(req.Password, user.PasswordHash) {
		resp := models.LoginSuccess{
			Message: "login successful",
			Token:   "1234-567-891011",
			User: models.UserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
		}
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
	}
}

func (a *AuthPassthroughHandler) SignUp(c *gin.Context) {
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

	if err := a.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	utils.PrettyPrint(user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func (a *AuthPassthroughHandler) comparePassword(pw string, hash string) bool {
	if hash == pw {
		return true
	} else {
		return false
	}
}

func hash(pw string) string {
	return pw
}
