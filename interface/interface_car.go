package interfaces

import "projectstructuring/models"

type ServiceCar interface {
	ServiceGetCar() (result models.Response)
}
