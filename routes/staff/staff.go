package staff

import (
	"github.com/gin-gonic/gin"
	"github.com/maxia51/bdgo/model"
	"github.com/maxia51/bdgo/repository"
)

type service struct {
	staffRepository repository.IStaffRepo
	roleRepository  repository.IRoleRepo
}

// New instanciate a staff router
// It return a staff pointer
func New(s repository.IStaffRepo, r repository.IRoleRepo) *service {
	return &service{
		staffRepository: s,
		roleRepository:  r,
	}
}

// Register add all the routes of the router
func (s *service) Register(r *gin.RouterGroup) {
	r.GET("/", s.getAllStaffHandler)
	r.POST("/", s.createStaffHandler)
	r.PUT("/", s.updtadeStaffHandler)
}

func (s *service) getAllStaffHandler(c *gin.Context) {

	staff, err := s.staffRepository.GetAll()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, staff)
}

func (s *service) createStaffHandler(c *gin.Context) {
	var staff model.Staff

	err := c.BindJSON(&staff)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = s.staffRepository.InsertStaff(&staff)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, staff)
}

func (s *service) updtadeStaffHandler(c *gin.Context) {

}
