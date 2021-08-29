package handler
import (
	"github.com/gofiber/fiber/v2"
	"homework-boiler-plate/models"
)

var(
	Cars = make([]models.Car,0)
)

func HandlerCreateCar(c *fiber.Ctx)error{
	body := new(models.Car)
	err := c.BodyParser(&body)

	if err != nil{
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": false,
			"message":"Cannot parse JSON",
			"error": err,
		})
	}

	if body.NameCar ==""|| body.ColorCar==""||body.TypeCar==""||body.MerkCar==""||body.Country==""||body.Engine==""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"message":"name, color, type,merk,country and engine are required",
		})
	}

	Cars = append(Cars,*body)

	if len(Cars) == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Fail Add New Car",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"status":  "OK",
		"data":    body,
	})

}

func GetAllCarsc(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"status":  "OK",
		"data":    Cars,
	})
}



