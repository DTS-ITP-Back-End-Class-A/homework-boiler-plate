package interfaces

import "tugasgolang/models"

type Service interface {
	ServiceGetCar() (result models.Response)
}
