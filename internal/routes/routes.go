package routes

import "github.com/gin-gonic/gin"

type UserInterface interface {
	GetUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
}

func RegisterRoutes(r *gin.Engine, users UserInterface) {
	r.GET("/users/:id", users.GetUserByID)
	r.POST("/users", users.CreateUser)

	// r.GET("/reviews", controllers.GetReviews)
	// r.POST("/reviews", controllers.CreateReview)

	// r.GET("/locations", controllers.GetLocations)
	// r.POST("/locations", controllers.CreateLocation)
}
