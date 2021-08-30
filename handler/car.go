package handler

import (
	"github.com/gofiber/fiber/v2"
)

func HandlerGetCar(c *fiber.Ctx) error {
	return c.SendString("masuk handler")
}
