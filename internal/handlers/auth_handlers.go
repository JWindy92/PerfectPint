package handlers

import "github.com/gin-gonic/gin"

/* -------------------------------------------------------------------------- */
/*                    Basic Auth Passthrough Implementation                   */
/* -------------------------------------------------------------------------- */

type AuthPassthroughHandler struct{}

func NewAuthPassthroughHandler() *AuthPassthroughHandler {
	return &AuthPassthroughHandler{}
}

func (a *AuthPassthroughHandler) Login(c *gin.Context) {

}

func (a *AuthPassthroughHandler) SignUp(c *gin.Context) {

}
