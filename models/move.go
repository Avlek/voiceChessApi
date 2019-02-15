package models

type MoveObject struct {
	Player Color  `json:"player"`
	Move   string `json:"move"`
}

type MoveObjects []MoveObject
