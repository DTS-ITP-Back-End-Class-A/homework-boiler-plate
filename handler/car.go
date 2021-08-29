package handler

import (
	//"projectstructuring/services"

	"github.com/gofiber/fiber/v2"
)

func HandlerGetCar(c *fiber.Ctx) error {
	//services.ServiceGetCar()
	return c.SendString("Masuk handler")

}
