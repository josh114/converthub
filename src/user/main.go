package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/user/database"
	"github.com/josh114/converthub/src/user/router"
)


func main() {
	database.ConnectDb()
	app := fiber.New()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":9015"))
}