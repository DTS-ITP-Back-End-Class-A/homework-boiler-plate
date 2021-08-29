package routes
import (
	"fmt"
	"homework-boiler-plate/handler"
	"github.com/gofiber/fiber/v2"
)


func Route(route fiber.Router){
fmt.Println("masuk route nih")
route.Get("/get",handler.GetAllCarsc)
route.Post("/create",handler.HandlerCreateCar)

}