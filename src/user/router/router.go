package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/user/api"
)


func SetupRoutes(app *fiber.App) {
	app.Post("/users", api.CreateUser)
	app.Get("/users", api.GetUsers)
}