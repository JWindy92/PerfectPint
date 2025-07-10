package main

import (
	"log"
	"net/http"

	"github.com/JWindy92/PerfectPint/internal/database"
	"github.com/JWindy92/PerfectPint/internal/handlers"
	"github.com/JWindy92/PerfectPint/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := database.PostgresImpl{}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	routes.RegisterRoutes(
		router,
		handlers.NewAuthPassthroughHandler(&db),
		handlers.NewDefaultUserHandler(&db),
		handlers.NewDefaultReviewHandler(&db),
		handlers.NewDefaultLocationHandler(&db),
	)

	router.Run("localhost:8080") //TODO: add config files
}

func index(c *gin.Context) {

	response := Response{Message: "Congrats! You made it!"}

	c.IndentedJSON(http.StatusOK, response)
}
