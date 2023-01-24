package main

import (
	"log"
	"net/http"

	"github.com/JWindy92/PerfectPint/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn := db.GetDatabase()
	defer db.CloseDatabase(conn)

	router := gin.Default()
	router.GET("/", index)

	router.Run("localhost:8080")
}

func index(c *gin.Context) {

	response := Response{Message: "Congrats! You made it!"}

	c.IndentedJSON(http.StatusOK, response)
}
