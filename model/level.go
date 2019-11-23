package model

type Level struct {
	Id    uint   `json:"id"`
	Level string `json:"level"`
}

type Levels []Level
