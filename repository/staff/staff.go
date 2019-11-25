package staff

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

func (s *service) GetStaffByEmail(email string) (model.Staff, error) {

	var staff model.Staff

	err := s.db.GetDatabase().QueryRow("SELECT staff.id, staff.email, staff.password, staff.created_at, staff.updated_at, role.* FROM staff INNER JOIN role ON staff.role = role.id WHERE email=?", email).Scan(&staff.Id, &staff.Email, &staff.Password, &staff.Created_at, &staff.Updated_at, &staff.Role.Id, &staff.Role.Name)

	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return staff, err
	}

	return staff, nil
}
