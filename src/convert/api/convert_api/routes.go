package convertapi

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func StartConvert(c *fiber.Ctx) error {
	var body Body
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).SendString(fmt.Sprintf("error occured while parsing body: %s", err))
	}
	fmt.Println(body.Filename, body.Format)
	 c.Status(200).SendString("successful")
	return nil
}