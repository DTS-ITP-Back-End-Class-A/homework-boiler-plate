package models

type Request struct {
	NameCar   string `json:"car_name"`
	ColorCar  string `json:"color_car"`
	TypeCar   string `json:"type_car"`
	MerkCar   string `json:"merk_car"`
	Country   string `json:"country"`
	Engine    string `json:"engine"`
}