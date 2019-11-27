package role

import (
	"github.com/maxia51/bdgo/database"
	"fmt"
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

func (s *service) GetRoleByID(id uint) (model.Role, error) {
	// TODO
	return model.Role{}, nil
}

func (s *service) GetRoleByName(name string) (model.Role, error) {

	var role model.Role

	row := s.db.GetDatabase().QueryRow("SELECT * FROM role WHERE name=? LIMIT 1", name)
	err := row.Scan(&role.Id, &role.Name)

	if err != nil || role.Name == "" {
		return role, fmt.Errorf("invalid role")
	}

	return role, nil
}
