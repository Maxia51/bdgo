package repository

import "github.com/maxia51/bdgo/model"

type IStaffRepo interface {
	GetStaffByEmail(email string) (model.Staff, error)
}
