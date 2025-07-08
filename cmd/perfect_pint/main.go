package main

import (
	"log"
	"net/http"

	"github.com/JWindy92/PerfectPint/internal/database"
	"github.com/JWindy92/PerfectPint/internal/handlers"
	"github.com/JWindy92/PerfectPint/internal/routes"
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
	db := database.SQLiteImpl{}

	router := gin.Default()

	routes.RegisterRoutes(
		router,
		handlers.NewAuthPassthroughHandler(&db),
		handlers.NewDefaultUserHandler(&db),
	)

	router.Run("localhost:8080") //TODO: add config files
}

func index(c *gin.Context) {

	response := Response{Message: "Congrats! You made it!"}

	c.IndentedJSON(http.StatusOK, response)
}
