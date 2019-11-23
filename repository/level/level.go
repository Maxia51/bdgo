package level

import (
	"fmt"
	"log"

	"github.com/maxia51/bdgo/database"
	"github.com/maxia51/bdgo/model"
)

type service struct {
	db database.IDatabase
}

func New(db database.IDatabase) *service {
	return &service{
		db: db,
	}
}

func (s *service) GetAll() (model.Levels, error) {

	var levels model.Levels

	results, err := s.db.GetDatabase().Query("SELECT * FROM level")

	if err != nil {
		return levels, err
	}

	for results.Next() {
		var level model.Level

		// for each row, scan the result into our tag composite object
		err = results.Scan(&level.Id, &level.Level)

		levels = append(levels, level)

		if err != nil {
			log.Panicf(err.Error()) // proper error handling instead of panic in your app
		}
	}

	if levels == nil {
		return levels, fmt.Errorf("Empty Levels")
	}

	return levels, nil

}
