package login

import (
	"net/http"

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

// New instanciate a login router
// It return a login pointer
func New(staff repository.IStaffRepo, security security.ISecurity) *router {
	return &router{
		staff:    staff,
		security: security,
	}
}

// Register add all the routes of the router
func (r *router) Register(gin *gin.RouterGroup) {
	gin.POST("/login", r.loginHandler)
}

// loginHandler handle the login path and create the jwt
func (r *router) loginHandler(c *gin.Context) {

	var json login

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := r.security.Auth(json.Email, json.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{"token": token})

}
