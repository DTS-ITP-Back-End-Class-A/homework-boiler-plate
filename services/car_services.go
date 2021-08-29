package services

import (
	"homework-boiler-plate/models"
	repo "homework-boiler-plate/repositories"
)

func ServicesGetCar() (result models.Car) {


	repo.RepoGetCar()
	return result
}