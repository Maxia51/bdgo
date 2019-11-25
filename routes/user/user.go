package user

import (
	"github.com/gin-gonic/gin"
	"github.com/maxia51/bdgo/security"
)

type router struct {
	security security.ISecurity
}

// New instanciate a user router
// It return a router pointer
func New(s security.ISecurity) *router {
	return &router{
		security: s,
	}
}

// Register add all the routes of the router
func (r *router) Register(gin *gin.RouterGroup) {
	gin.GET("/user", r.userHandler)
}

// userHandler handle the user request 
func (r *router) userHandler(c *gin.Context) {

	// TODO implement correct response

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
