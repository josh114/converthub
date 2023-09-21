package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/upload_api/api"
)


func SetupRoutes(app *fiber.App) {
	app.Post("/upload", api.Upload)

}