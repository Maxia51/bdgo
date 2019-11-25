package login

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maxia51/bdgo/repository"
	"github.com/maxia51/bdgo/security"
)

type router struct {
	staff    repository.IStaffRepo
	security security.ISecurity
}

type login struct {
	Email    string `form:"email" json:"email" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func New(staff repository.IStaffRepo, security security.ISecurity) *router {
	return &router{
		staff:    staff,
		security: security,
	}
}

func (r *router) Register(gin *gin.RouterGroup) {
	gin.POST("/login", r.loginHandler)
}

func (r *router) loginHandler(c *gin.Context) {

	var json login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := r.security.Auth(json.Email, json.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	fmt.Println(r.security.IsLogged())
	fmt.Println(r.security.GetSession())

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

}
