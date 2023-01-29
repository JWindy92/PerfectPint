package api

import (
	"log"
	"net/http"

	"github.com/JWindy92/PerfectPint/pkg/auth"
	"github.com/JWindy92/PerfectPint/pkg/db"
	"github.com/JWindy92/PerfectPint/pkg/models"
	"github.com/gin-gonic/gin"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(c *gin.Context) {

	var new_user models.User
	c.BindJSON(&new_user)

	var conn = db.GetDatabase()
	defer db.CloseDatabase(conn)

	var db_user models.User
	conn.Where("email = ?", new_user.Email).First(&db_user)

	if db_user.Email != "" {
		var err models.Error
		err = models.SetError(err, "Email already in use")

		c.IndentedJSON(http.StatusOK, err) //TODO: return different error code?
		return
	}

	password_hash, err := auth.GeneratehashPassword(new_user.Password)
	new_user.Password = password_hash
	if err != nil {
		log.Fatalln("error creating password hash")
	}

	conn.Create(&new_user)
	c.IndentedJSON(http.StatusCreated, new_user)
}

func Login(c *gin.Context) {
	log.Println("Logging in")
}
