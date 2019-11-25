package model

import "time"

type Staff struct {
	Id         uint
	Email      string
	Password   string
	Role       Role
	Created_at time.Time
	Updated_at time.Time
}

type Staffs []Staffs