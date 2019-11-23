package main

import (
	"fmt"

	"github.com/maxia51/bdgo/database"
	levelRepository "github.com/maxia51/bdgo/repository/level"
	userRepository "github.com/maxia51/bdgo/repository/user"
)

func main() {

	db := database.New("root", "example", "127.0.0.1", "3306", "bde")

	levelRepository := levelRepository.New(db)
	userRepository := userRepository.New(db)



	fmt.Println(levelRepository.GetAll())
	fmt.Println(userRepository.GetAll())
	fmt.Println("hello world")
}
