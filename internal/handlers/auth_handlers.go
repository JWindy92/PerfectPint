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
	name := c.Query("name")
	email := c.Query("email")
	user := models.User{
		Name:  name,
		Email: email,
	}
	a.UserService.CreateUser(&user)
}

func (a *AuthPassthroughHandler) comparePassword(pw string, hash string) bool {
	if hash == pw {
		return true
	} else {
		return false
	}
}
