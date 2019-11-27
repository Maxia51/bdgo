package model

import "time"

// Staff DTO
type Staff struct {
	Id         uint 		`json:"id"`
	Email      string		`json:"email"`
	Password   string		`json:"password,omitempty"`
	Role       Role			`json:"role"`
	Created_at time.Time	`json:"created_at"`
	Updated_at time.Time	`json:"updated_at"`
}

// Staffs DTO
type Staffs []Staff 