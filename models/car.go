package models

import "time"

type Car struct {
	NameCar   string
	ColorCar  string
	TypeCar   string
	MerkCar   string
	Country   string
	Engine    string
	CreatedAt time.Time
	UpadateAt time.Time
}
