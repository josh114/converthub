package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/auth/api"
)

func welcome(c *fiber.Ctx) error {
	c.SendString("welcome to my request api")
	return nil
}


func SetupRoutes(app *fiber.App) {
	app.Get("/auth", welcome)
	app.Post("/auth", api.Login)
}