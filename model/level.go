package model

// Level DTO
type Level struct {
	Id    uint   `json:"id"`
	Level string `json:"level"`
}

// Levels DTO
type Levels []Level
