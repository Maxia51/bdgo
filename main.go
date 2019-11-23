package main

import (
	"fmt"
	"os"

	"github.com/maxia51/bdgo/database"
	levelRepository "github.com/maxia51/bdgo/repository/level"
	userRepository "github.com/maxia51/bdgo/repository/user"
)

func main() {

	// for dev purpose
	os.Setenv("MYSQL_USER", "root")
	os.Setenv("MYSQL_USER_PASSWORD", "example")
	os.Setenv("MYSQL_ADDR", "127.0.0.1")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_DATABASE_NAME", "bde")

	db := database.New()

	levelRepository := levelRepository.New(db)
	userRepository := userRepository.New(db)

	fmt.Println(levelRepository.GetAll())
	fmt.Println(userRepository.GetAll())
	fmt.Println("hello world")
}
