package staff

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/maxia51/bdgo/database"
	"github.com/maxia51/bdgo/repository"
	"github.com/maxia51/bdgo/model"
)

type service struct {
	db database.IDatabase
	roleRepository repository.IRoleRepo
}

// New instanciate a staff repository
// Return a pointer of the staff repository
func New(db database.IDatabase, role repository.IRoleRepo) *service {
	return &service{
		db: db,
		roleRepository: role,
	}
}

// GetAll return all the staff of the app
func (s *service) GetAll() (model.Staffs, error) {

	var staffs model.Staffs

	result, err := s.db.GetDatabase().Query("SELECT staff.id, email, staff.role, role.name, created_at, updated_at FROM staff INNER JOIN role ON staff.role = role.id")

	if err != nil {
		return staffs, err
	}

	for result.Next() {

		var staff model.Staff

		result.Scan(&staff.Id, &staff.Email, &staff.Role.Id, &staff.Role.Name, &staff.Created_at, &staff.Updated_at)

		staffs = append(staffs, staff)

		if err != nil {
			log.Panicf(err.Error()) // proper error handling instead of panic in your app
		}

	}

	return staffs, nil
}

// GetStaffByEmail get a staff member by email
// It return a model.staff struct and error
func (s *service) GetStaffByEmail(email string) (model.Staff, error) {

	var staff model.Staff

	// scan and populate staff struct
	err := s.db.GetDatabase().QueryRow("SELECT staff.id, staff.email, staff.password, staff.created_at, staff.updated_at, role.* FROM staff INNER JOIN role ON staff.role = role.id WHERE email=?", email).Scan(&staff.Id, &staff.Email, &staff.Password, &staff.Created_at, &staff.Updated_at, &staff.Role.Id, &staff.Role.Name)

	if err != nil {
		return staff, err
	}

	return staff, nil
}

// InsertStaff Create a new Staff member
// It return the staff pointer created and error
func (s *service) InsertStaff(staff *model.Staff) error {

	// Check email
	err := s.emailValidator(staff.Email)

	if err != nil {
		return fmt.Errorf("email invalid")
	}

	// Check role

	role, err := s.roleRepository.GetRoleByName(staff.Role.Name)

	if err != nil {
		return err
	}

	// Starting insert

	creationTime := time.Now()
	dbStaff, err := s.db.GetDatabase().Exec("INSERT INTO staff(email, password, role, created_at, updated_at) VALUES (?,?,?,?,?)", staff.Email, staff.Password, role.Id, creationTime, creationTime)

	if err != nil {
		return err
	}

	// Populate struct
	dbID, err := dbStaff.LastInsertId()

	if err != nil {
		return err
	}

	staff.Id = uint(dbID)
	staff.Password = ""
	staff.Role.Id = role.Id
	staff.Role.Name = role.Name
	staff.Created_at = creationTime
	staff.Updated_at = creationTime

	return nil
}

// UpdateStaff update a staff member
func (s *service) UpdateStaff(staff *model.Staff) (error) {
	// TODO
	return nil
}

// emailValidator check if email as correct format
func (s *service) emailValidator(email string) (error) {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) > 254 || !rxEmail.MatchString(email) {
		return fmt.Errorf("error: foo is not a valid email address")
	}

	return nil
}
