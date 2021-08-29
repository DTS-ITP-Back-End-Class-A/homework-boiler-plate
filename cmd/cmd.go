package cmd

import (
	"fmt"
	"homework-boiler-plate/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)



func Run() {
	fmt.Println("masukk cmd nih")
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)

}

func setupRoutes(app *fiber.App) {

	api := app.Group("/")
	routes.Route(api)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}