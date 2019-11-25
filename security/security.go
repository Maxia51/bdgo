package security

import "github.com/maxia51/bdgo/repository"

import "fmt"

type service struct {
	db repository.IStaffRepo
	Session Session
}

type Session struct {
	Token string
	Role string
}

func New(db repository.IStaffRepo) *service {

	return &service{
		db: db,
		Session: Session{
			Token: "",
			Role: "",
		},
	}
}

func (s *service) IsLogged() bool {
	if len(s.Session.Token) > 0 {
		return true
	}
	return false
}

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

	s.Session.Token = "adding"
	s.Session.Role = "admin"

	return true, nil
}


func (s *service) GetSession() (Session, error) {
	if len(s.Session.Token) < 0 || len(s.Session.Role) < 0 {
		return s.Session, fmt.Errorf("Session not initialized")
	}

	return s.Session, nil
}
