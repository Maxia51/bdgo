package user

import (
	"github.com/gin-gonic/gin"
)

type router struct {
}

// New instanciate a user router
// It return a router pointer
func New() *router {
	return &router{}
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
