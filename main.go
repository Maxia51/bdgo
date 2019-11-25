package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/maxia51/bdgo/database"
	staffRepository "github.com/maxia51/bdgo/repository/staff"
	"github.com/maxia51/bdgo/routes/user"
	"github.com/maxia51/bdgo/routes/login"
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
	staffRepository := staffRepository.New(db)

	securityService := security.New(staffRepository)

	fmt.Println("hello world")

	router := gin.Default()

	api := router.Group("/api")
	{

		loginRouter := login.New(staffRepository, securityService)
		loginRouter.Register(api)
	}
	{
		adminAuth := api.Group("/v1")
		userRouter := user.New(securityService)
		userRouter.Register(adminAuth)
	}

	router.Run(":3000")
}
