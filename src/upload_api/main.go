package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/upload_api/database"
)


func main () {
	database.ConnectDb()
	app:= fiber.New()

	log.Fatal(app.Listen(":9010"))
}