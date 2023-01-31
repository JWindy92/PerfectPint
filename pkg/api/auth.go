package api

import (
	"log"
	"net/http"

	"github.com/JWindy92/PerfectPint/pkg/auth"
	"github.com/JWindy92/PerfectPint/pkg/db"
	"github.com/JWindy92/PerfectPint/pkg/models"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}
type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
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
	var incoming_auth Authentication
	c.BindJSON(&incoming_auth)

	var conn = db.GetDatabase()
	defer db.CloseDatabase(conn)

	var db_user models.User
	conn.Where("email = ?", incoming_auth.Email).First(&db_user)
	if db_user.Email == "" {
		var err models.Error
		err = models.SetError(err, "Username or Password is incorrect")

		c.IndentedJSON(http.StatusUnauthorized, err)
		return
	}

	check := auth.CheckPasswordHash(incoming_auth.Password, db_user.Password)

	if !check {
		var err models.Error
		err = models.SetError(err, "Username or Password is incorrect")

		c.IndentedJSON(http.StatusUnauthorized, err)
		return
	}

	//TODO: generate JWT and return
	validToken, err := auth.GenerateJWT(db_user.Email, db_user.Role)
	if err != nil {
		var err models.Error
		err = models.SetError(err, "Failed to generate token")

		c.IndentedJSON(http.StatusUnauthorized, err)
		return
	}

	var token Token
	token.Email = db_user.Email
	token.Role = db_user.Role
	token.TokenString = validToken

	c.IndentedJSON(http.StatusOK, token)
}
