package repository

import "github.com/maxia51/bdgo/model"

// IStaffRepo is the interface that wraps the basic staff repository method.
//
// GetStaffByEmail get the staff member by email.
// It returns a model.staff struct.
// and any error encountered that caused the get fail.
type IStaffRepo interface {
	GetAll() (model.Staffs, error)
	GetStaffByEmail(email string) (model.Staff, error)
	InsertStaff(staff *model.Staff) (error)
	UpdateStaff(staff *model.Staff) (error)
}

type IRoleRepo interface {
	GetRoleByID(id uint) (model.Role, error)
	GetRoleByName(name string) (model.Role, error)
}