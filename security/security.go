package security

import "github.com/maxia51/bdgo/repository"

import "fmt"

type service struct {
	db repository.IStaffRepo
}

// New instanciate a security service
// Security service implement ISecurity interface
// It return a security pointer
func New(db repository.IStaffRepo) *service {

	return &service{
		db: db,
	}
}

// Auth check the posted credentail
// It return a boolean value and error
func (s *service) Auth(email string, password string) (bool, error) {

	staff, err := s.db.GetStaffByEmail(email)

	fmt.Println(staff)

	if err != nil {
		return false, err
	}

	// Check password
	if staff.Password != password {
		return false, fmt.Errorf("Incorrect Password")
	}

	return true, nil
}