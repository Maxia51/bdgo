package repository

import "github.com/maxia51/bdgo/model"

// IStaffRepo is the interface that wraps the basic staff repository method.
//
// GetStaffByEmail get the staff member by email.
// It returns a model.staff struct.
// and any error encountered that caused the get fail.
type IStaffRepo interface {
	GetStaffByEmail(email string) (model.Staff, error)
}
