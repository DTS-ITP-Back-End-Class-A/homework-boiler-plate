package services

import (
	"fmt"
	"tugasgolang/models"
	"tugasgolang/repositories"
)

// func Service() {
// 	fmt.Println("masuk service")
// }

func ServiceGetCar() (result models.Response) {
	fmt.Println("masuk Service")
	repositories.RepoGetCar()
	return result
}
