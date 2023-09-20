package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/auth/database"
	"github.com/josh114/converthub/src/auth/router"
)


func main () {
	database.ConnectDb()
	app := fiber.New()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":9005"))
}