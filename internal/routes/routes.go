package routes

import (
	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
}

type UserInterface interface {
	GetUserByID(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	CreateUser(c *gin.Context)
}

func RegisterRoutes(
	r *gin.Engine,
	auth AuthInterface,
	users UserInterface,
) {

	r.POST("/signup", auth.SignUp)
	r.POST("/login", auth.Login)

	r.GET("/users/:id", users.GetUserByID)
	r.GET("/users", users.GetUserByEmail) //TODO: should make less ambiguous. Maybe implement a broader search function
	r.POST("/users", users.CreateUser)

	// r.GET("/reviews", controllers.GetReviews)
	// r.POST("/reviews", controllers.CreateReview)

	// r.GET("/locations", controllers.GetLocations)
	// r.POST("/locations", controllers.CreateLocation)
}
