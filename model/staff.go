package model

import "time"

// Staff DTO
type Staff struct {
	Id         uint
	Email      string
	Password   string
	Role       Role
	Created_at time.Time
	Updated_at time.Time
}

// Staffs DTO
type Staffs []Staffs