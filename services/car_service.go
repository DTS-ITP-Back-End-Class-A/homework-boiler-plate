package services

import (
	"projectstructuring/models"
	repo "projectstructuring/repositories"
)

// func ServiceGetCar(c *fiber.Ctx) error {
// 	return c.SendString("Masuk service")
// }

func ServiceGetCar() (result models.Response) {
	repo.RepoGetCar()
	return result
}
