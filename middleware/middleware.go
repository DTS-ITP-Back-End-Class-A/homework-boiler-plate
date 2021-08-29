package middleware
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuthMiddleware() func(*fiber.Ctx) error {
	basicAuthHandler := basicauth.New(basicauth.Config{
		Users : map[string]string{
			"admin":"12345",
		},
		Realm:"Forbidden",
		Authorizer: func(username, password string)bool{
			if username =="admin"&& password == "12345"{
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message":"User Doesn't have access",
				"error":"Unauthorised",
			})
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	})

	return basicAuthHandler
}