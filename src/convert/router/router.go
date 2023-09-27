package router

import (
	"github.com/gofiber/fiber/v2"
	convertapi "github.com/josh114/converthub/src/convert/api/convert_api"
	uploadapi "github.com/josh114/converthub/src/convert/api/upload_api"
)


func SetupRoutes(app *fiber.App) {
	app.Post("/upload", uploadapi.UploadFile)
	app.Post("/convert", convertapi.StartConvert)
}