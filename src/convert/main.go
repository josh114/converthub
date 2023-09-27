package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/convert/database"
	"github.com/josh114/converthub/src/convert/router"
)


func main () {
	database.ConnectDb()
	app:= fiber.New(fiber.Config{
		BodyLimit: 1000 * 1024 *1024,
	})
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":9020"))
}