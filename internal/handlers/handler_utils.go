package handlers

import (
	"github.com/JWindy92/PerfectPint/internal/database"
	"github.com/JWindy92/PerfectPint/internal/routes"
	"github.com/gin-gonic/gin"
)

func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	db := database.SQLiteImpl{}

	// uHandler := NewDefaultUserHandler(&db)
	r := gin.Default()
	routes.RegisterRoutes(
		r,
		NewAuthPassthroughHandler(&db),
		NewDefaultUserHandler(&db),
		NewDefaultReviewHandler(&db),
		NewDefaultLocationHandler(&db),
	)
	return r
}
