package handlers

import (
	"github.com/JWindy92/PerfectPint/internal/database"
	"github.com/JWindy92/PerfectPint/internal/models"
	"github.com/JWindy92/PerfectPint/internal/services"
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

func NewAuthPassthroughHandler() *AuthPassthroughHandler {
	return &AuthPassthroughHandler{
		UserService: services.NewUserService(database.DB),
	}
}

func (a *AuthPassthroughHandler) Login(c *gin.Context) {
	email := c.Query("email")
	// pw := c.Query("password")

	a.UserService.GetUserByEmail(email)
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
