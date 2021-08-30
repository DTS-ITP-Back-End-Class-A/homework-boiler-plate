package routes

import (
	"fmt"
	"tugasgolang/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes() {
	fmt.Println("masuk route")
	app := fiber.New()
	// app.Get("/get-car", func(c *fiber.Ctx) error {
	// 	return c.SendString("fiberjalan")
	// })
	app.Get("/get-car", handler.HandlerGetCar)
	app.Listen(":3000")
}
