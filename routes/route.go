package routes

import (
	"fmt"
	"projectstructuring/handler"

	// "projectstructuring/repositories"
	// "projectstructuring/services"

	"github.com/gofiber/fiber/v2"
)

func Route() {
	fmt.Println("Masuk route")

	app := fiber.New()

	app.Get("/get-car", handler.HandlerGetCar)
	// app.Get("/get-car1", services.ServiceGetCar)
	// app.Get("/get-car2", repositories.RepoCar)

	app.Listen(":3000")
}
