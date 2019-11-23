package model

import (
	"time"
)

type User struct {
	Id         uint
	Firstname  string
	Lastname   string
	Money      float32
	Level    Level
	Created_at time.Time
	Updated_at time.Time
}

type Users []User
