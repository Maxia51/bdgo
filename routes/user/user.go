package user

import (
	"github.com/gin-gonic/gin"
	"github.com/maxia51/bdgo/security"
)

type router struct {
	security security.ISecurity
}

func New(s security.ISecurity) *router {
	return &router{
		security: s,
	}
}

func (r *router) Register(gin *gin.RouterGroup) {
	gin.GET("/user", r.userHandler)
}

func (r *router) userHandler(c *gin.Context) {

	session, _ := r.security.GetSession()

	c.JSON(200, gin.H{
		"isLogged": r.security.IsLogged(),
		"token": session,
		"message": "pong",
	})
}
