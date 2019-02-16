package models

type MoveObject struct {
	Player uint8  `json:"user"`
	Move   string `json:"move"`
}

type MoveObjects []MoveObject
