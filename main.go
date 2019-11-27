package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/maxia51/bdgo/database"
	"github.com/maxia51/bdgo/middleware"
	staffRepository "github.com/maxia51/bdgo/repository/staff"
	roleRepository "github.com/maxia51/bdgo/repository/role"
	"github.com/maxia51/bdgo/routes/login"
	"github.com/maxia51/bdgo/routes/staff"
	"github.com/maxia51/bdgo/security"
)

func main() {

	// for dev purpose
	os.Setenv("MYSQL_USER", "root")
	os.Setenv("MYSQL_USER_PASSWORD", "example")
	os.Setenv("MYSQL_ADDR", "127.0.0.1")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_DATABASE_NAME", "bde")
	os.Setenv("JWT_SECRET_KEY", "bde")

	db := database.New()

	// levelRepository := levelRepository.New(db)
	// userRepository := userRepository.New(db)
	roleRepository := roleRepository.New(db)
	staffRepository := staffRepository.New(db, roleRepository)

	securityService := security.New(staffRepository)

	fmt.Println("hello world")

	router := gin.Default()

	api := router.Group("/api/v1")
	{

		loginRouter := login.New(staffRepository, securityService)
		loginRouter.Register(api)
	}
	{
		// ADMIN Routes
		adminRouter := api.Group("/staff")
		{
			adminRouter.Use(middleware.AuthRequired("ADMIN"))
			{
				staffRouter := staff.New(staffRepository, roleRepository)
				staffRouter.Register(adminRouter)
			}
		}
	}
	{
		// MODERATOR Routes
		//moderatorRouter := api.Group("/user")
	}

	router.Run(":3000")
}
