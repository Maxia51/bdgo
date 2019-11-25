package user

import (
	"log"
	"fmt"

	"github.com/maxia51/bdgo/database"
	"github.com/maxia51/bdgo/model"
)

type service struct {
	db database.IDatabase
}

// New Instanciate a user repository
// It return a user pointer
func New(db database.IDatabase) *service {
	return &service{
		db: db,
	}
}

// GetAll get all the users 
// It return models.Users and error
func (s *service) GetAll() (model.Users, error) {

	var users model.Users

	results, err := s.db.GetDatabase().Query("SELECT user.*, level.level FROM user INNER JOIN level ON user.level = level.id")

	if err != nil {
		return users, err
	}

	for results.Next() {

		var user model.User

		// populate user struct
		err = results.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Money, &user.Level.Id, &user.Created_at, &user.Updated_at, &user.Level.Level)

		users = append(users, user)

		if err != nil {
			log.Panicf(err.Error()) // proper error handling instead of panic in your app
		}

	}

	if users == nil {
		return users, fmt.Errorf("Empty users")
	} 

	return users, nil

}
