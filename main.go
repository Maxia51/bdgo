package main

import (
	"fmt"

	"github.com/maxia51/bdgo/database"
)

func main() {

	database.New("root", "example", "127.0.0.1", "3306", "bde")

	fmt.Println("hello world")
}
